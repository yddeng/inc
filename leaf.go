package inc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	net2 "github.com/yddeng/inc/net"
	"github.com/yddeng/utils/task"
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
	dialers map[string]*dialer

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
		dialers:   map[string]*dialer{},
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

	dialer := &dialer{tunID: id, address: msg.GetAddress(), leaf: this}
	//if err := leaf.dial(); err != nil {
	//	replier.Reply(&net2.CreateTunnelResp{Msg: err.Error()}, nil)
	//	return
	//}

	this.counter++
	this.dialers[id] = dialer
	replier.Reply(&net2.CreateTunnelResp{TunnelID: id}, nil)
}

func (this *Leaf) onTunnelMessage(replier *drpc.Replier, req interface{}) {
	msg := req.(*net2.TunnelMessageReq)
	fmt.Println("onTunnelMessage", msg)
	dialer, ok := this.dialers[msg.GetTunID()]
	if !ok {
		replier.Reply(&net2.TunnelMessageResp{Msg: "tunnel is not exist"}, nil)
		return
	}

	if msg.GetClose() {
		dialer.close()
	} else {
		dialer.connID = msg.GetConnID()
		if err := dialer.write(msg.GetData()); err != nil {
			replier.Reply(&net2.TunnelMessageResp{Msg: "conn.Write " + err.Error()}, nil)
			return
		}
	}
	replier.Reply(&net2.TunnelMessageResp{}, nil)
}

type dialer struct {
	leaf    *Leaf
	tunID   string
	address string
	connID  uint32
	conn    net.Conn
}

func (this *dialer) close() {
	if this.conn != nil {
		this.conn.Close()
		this.conn = nil
	}
}

func (this *dialer) write(b []byte) (err error) {
	if this.conn == nil {
		if this.conn, err = net.Dial("tcp", this.address); err != nil {
			return
		}
		go this.handleConn()
	}
	if _, err = this.conn.Write(b); err != nil {
		return
	}
	return
}

func (this *dialer) handleConn() {
	go func() {
		buf := make([]byte, 1024)
		for {
			msg := &net2.TunnelMessageReq{TunID: this.tunID, ConnID: this.connID}

			n, err := this.conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Read", err)
				msg.Close = true
				_, _ = this.leaf.rpcClient.Call(this.leaf, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout)
				break
			}

			fmt.Println("conn.Read", buf[:n])
			msg.Data = buf[:n]
			if _, err := this.leaf.rpcClient.Call(this.leaf, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout); err != nil {
				break
			}
		}

		this.leaf.taskQueue.Push(func() {
			this.close()
		})

	}()
}
