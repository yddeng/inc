package intun

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/dutil/task"
	net2 "github.com/yddeng/intun/net"
	"net"
	"reflect"
	"time"
)

type Leaf struct {
	id    uint32
	name  string
	rAddr string

	taskQueue *task.TaskQueue

	counter uint32
	tunnel  map[string]*leafTunnel

	dialing bool
	session dnet.Session

	rpcServer *drpc.Server
	rpcClient *drpc.Client
}

func (this *Leaf) SendRequest(req *drpc.Request) error {
	return this.session.Send(req)
}

func (this *Leaf) SendResponse(resp *drpc.Response) error {
	return this.session.Send(resp)
}

func (this *Leaf) dial() {
	if this.dialing {
		return
	}

	this.dialing = true

	go func() {
		for {
			if conn, err := dnet.DialTCP(this.rAddr, time.Second*5); nil == err && conn != nil {
				this.onConnected(conn)
				return
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (this *Leaf) onConnected(conn dnet.NetConn) {
	this.taskQueue.Push(func() {
		this.dialing = false

		this.session = dnet.NewTCPSession(conn,
			dnet.WithCodec(net2.NewCodec()),
			dnet.WithCloseCallback(func(session dnet.Session, reason error) {
				this.taskQueue.Push(func() {
					this.session = nil
					fmt.Printf("onConnected session closed, reason: %s\n", reason)
					this.dial()
				})
			}),
			dnet.WithErrorCallback(func(session dnet.Session, err error) {
				fmt.Println("onConnected session error:", err)
				session.Close(err)
			}),
			dnet.WithMessageCallback(func(session dnet.Session, data interface{}) {
				this.taskQueue.Push(func() {
					var err error
					switch data.(type) {
					case *drpc.Request:
						err = this.rpcServer.OnRPCRequest(&endpoint{session: session}, data.(*drpc.Request))
					case *drpc.Response:
						err = this.rpcClient.OnRPCResponse(data.(*drpc.Response))
					default:
						err = fmt.Errorf("invalid type:%s", reflect.TypeOf(data).String())
					}
					if err != nil {
						fmt.Printf("onConnected dispatch error: %s. \n", err.Error())
					}
				})
			}),
		)

		// 注册身份信息
		req := &net2.LeafRegisterReq{Name: this.name}
		if err := this.rpcClient.Go(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
			if e != nil || i.(*net2.LeafRegisterResp).GetMsg() != "" {
				var msg string
				if e != nil {
					msg = fmt.Sprintf("onConnected loginResp failed, error %s", e.Error())
				} else {
					msg = fmt.Sprintf("onConnected loginResp false, msg %s", i.(*net2.LeafRegisterResp).GetMsg())
				}
				fmt.Println(msg)
				panic(msg)
				return
			}

			fmt.Println("onConnected login center ok")
			msg := i.(*net2.LeafRegisterResp)
			this.id = msg.GetID()
		}); err != nil {

		}
	})
}

func LaunchLeaf(name, rootAddr string) *Leaf {
	taskQueue := task.NewTaskQueue(512)
	taskQueue.Run()

	leaf := &Leaf{
		name:      name,
		rAddr:     rootAddr,
		taskQueue: taskQueue,
		counter:   1,
		tunnel:    map[string]*leafTunnel{},
		rpcServer: drpc.NewServer(),
		rpcClient: drpc.NewClient(),
	}

	leaf.rpcServer.Register(proto.MessageName(&net2.CreateTunnelReq{}), leaf.onCreateTunnel)
	leaf.rpcServer.Register(proto.MessageName(&net2.TunnelMessageReq{}), leaf.onTunnelMessage)

	leaf.taskQueue.Push(func() { leaf.dial() })
	return leaf

}

func (this *Leaf) onCreateTunnel(replier *drpc.Replier, req interface{}) {
	msg := req.(*net2.CreateTunnelReq)
	fmt.Println("onCreateTunnel", msg)
	id := fmt.Sprintf("%d:%d", this.id, this.counter)

	leaf := &leafTunnel{tunID: id, address: msg.GetAddress(), leaf: this,
		buf: make([]byte, 1024)}
	//if err := leaf.dial(); err != nil {
	//	replier.Reply(&net2.CreateTunnelResp{Msg: err.Error()}, nil)
	//	return
	//}

	this.counter++
	this.tunnel[id] = leaf
	replier.Reply(&net2.CreateTunnelResp{TunnelID: id}, nil)
}

func (this *Leaf) onTunnelMessage(replier *drpc.Replier, req interface{}) {
	msg := req.(*net2.TunnelMessageReq)
	fmt.Println("onTunnelMessage", msg)
	tun, ok := this.tunnel[msg.GetTunID()]
	if !ok {
		replier.Reply(&net2.TunnelMessageResp{Msg: "tunnel is not exist"}, nil)
		return
	}

	if tun.conn == nil {
		tun.dial()
	}

	_, err := tun.conn.Write(msg.GetData())
	if err != nil {
		replier.Reply(&net2.TunnelMessageResp{Msg: "conn.Write " + err.Error()}, nil)
		return
	}
	replier.Reply(&net2.TunnelMessageResp{}, nil)
}

type leafTunnel struct {
	leaf    *Leaf
	tunID   string
	address string
	conn    net.Conn
	buf     []byte
}

func (this *leafTunnel) dial() error {
	conn, err := net.Dial("tcp", this.address)
	if err != nil {
		return err
	}
	this.conn = conn
	// read from conn
	go func() {
		for {
			n, err := conn.Read(this.buf)
			if err != nil {
				fmt.Println("conn.Read", err)
				break
			}

			msg := &net2.TunnelMessageReq{
				TunID: this.tunID,
				Data:  this.buf[:n],
			}
			fmt.Println("conn.Read", this.buf[:n])
			if _, err := this.leaf.rpcClient.Call(this.leaf, proto.MessageName(&net2.TunnelMessageReq{}), msg, drpc.DefaultRPCTimeout); err != nil {
				break
			}
		}

		conn.Close()
		this.leaf.taskQueue.Push(func() {
			this.conn = nil
		})

	}()
	return nil
}
