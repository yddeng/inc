package inc

/*
import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/inc/net"
	"github.com/yddeng/utils/task"
	net2 "net"
	"reflect"
	"time"
)

type Client struct {
	address   string
	password  string
	session   dnet.Session
	taskQueue *task.TaskQueue
	rpcServer *drpc.Server
	rpcClient *drpc.Client

	id       uint32
	leaf     uint32
	acceptor *acceptor
}

func (this *Client) SendRequest(req *drpc.Request) error {
	return this.session.Send(req)
}

func (this *Client) SendResponse(resp *drpc.Response) error {
	return this.session.Send(resp)
}

func LaunchClient(address, password string) *Client {
	conn, err := dnet.DialTCP(address, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	taskQueue := task.NewTaskQueue(128)
	taskQueue.Run()

	client := &Client{
		address:   address,
		password:  password,
		taskQueue: taskQueue,
		leaf:      1,
		rpcServer: drpc.NewServer(),
		rpcClient: drpc.NewClient(),
	}

	client.rpcServer.Register(proto.MessageName(&net.TunnelMessageReq{}), client.onTunnelMessage)

	client.session = dnet.NewTCPSession(conn,
		dnet.WithCodec(net.NewCodec()),
		dnet.WithCloseCallback(func(session dnet.Session, reason error) {
			client.taskQueue.Push(func() {
				client.session = nil
				fmt.Printf("onConnected session closed, reason: %s\n", reason)
			})
		}),
		dnet.WithErrorCallback(func(session dnet.Session, err error) {
			fmt.Println("onConnected session error:", err)
			session.Close(err)
		}),
		dnet.WithMessageCallback(func(session dnet.Session, data interface{}) {
			client.taskQueue.Push(func() {
				var err error
				switch data.(type) {
				case *drpc.Request:
					err = client.rpcServer.OnRPCRequest(client, data.(*drpc.Request))
				case *drpc.Response:
					err = client.rpcClient.OnRPCResponse(data.(*drpc.Response))
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
	req := &net.CliAuthReq{Password: password}
	ret, err := client.rpcClient.Call(client, proto.MessageName(req), req, drpc.DefaultRPCTimeout)
	if err != nil {
		fmt.Println(err)
		panic(err)

	}

	fmt.Println("onConnected login center ok")
	resp := ret.(*net.CliAuthResp)
	if resp.GetMsg() != "" {
		fmt.Println(err)
		panic(err)
	}
	client.id = resp.GetID()

	return client
}

func (this *Client) SelectLeaf(leaf uint32) {
	this.leaf = leaf
	fmt.Println("select leaf ok", leaf)
}

func (this *Client) CreateTunnel(laddr, raddr string) error {
	errC := make(chan error, 1)

	f := func() {
		if this.leaf == 0 {
			errC <- fmt.Errorf("select leaf before")
			return
		}

		req := &net.CreateTunnelReq{
			LeafID:  this.leaf,
			Address: raddr,
		}

		if err := this.rpcClient.Go(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
			resp := i.(*net.CreateTunnelResp)
			if resp.GetMsg() != "" {
				errC <- fmt.Errorf(resp.GetMsg())
				return
			}

			acceptor := &acceptor{
				client: this,
				tunID:  resp.GetTunnelID(),
				conns:  map[uint32]net2.Conn{},
			}

			if err := acceptor.listen(laddr); err != nil {
				errC <- err
				return
			}

			fmt.Println("create tunnel ok", resp.GetTunnelID())
			this.acceptor = acceptor
			errC <- nil
		}); err != nil {
			errC <- err
		}
	}

	this.taskQueue.Push(f)
	return <-errC
}

func (this *Client) onTunnelMessage(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.TunnelMessageReq)
	if this.acceptor == nil {
		replier.Reply(&net.TunnelMessageResp{Msg: "tunnel is not exist"}, nil)
		return
	}

	conn, ok := this.acceptor.conns[msg.GetConnID()]
	if !ok {
		replier.Reply(&net.TunnelMessageResp{Msg: "tunnel is not exist"}, nil)
		return
	}

	if msg.GetClose() {
		fmt.Println("ontunnel close")
		this.acceptor.close(msg.GetConnID())
	} else {
		fmt.Println("--", msg.GetData())
		_, err := conn.Write(msg.GetData())
		if err != nil {
			replier.Reply(&net.TunnelMessageResp{Msg: "conn.Write " + err.Error()}, nil)
			return
		}
	}
	replier.Reply(&net.TunnelMessageResp{}, nil)
}

type acceptor struct {
	client   *Client
	tunID    string
	listener net2.Listener
	counter  uint32
	conns    map[uint32]net2.Conn
}

func (this *acceptor) close(id uint32) {
	c, ok := this.conns[id]
	if ok {
		c.Close()
		delete(this.conns, id)
	}
}

func (this *acceptor) listen(addr string) (err error) {
	if this.listener, err = net2.Listen("tcp", addr); err != nil {
		return
	}

	go func() {
		for {
			conn, err := this.listener.Accept()
			if err != nil {
				if ne, ok := err.(net2.Error); ok && ne.Temporary() {
					continue
				} else {
					panic(err)
				}
			}
			this.handleConn(conn)
		}
	}()
	return
}

func (this *acceptor) handleConn(conn net2.Conn) {
	this.client.taskQueue.Push(func() {
		id := this.counter
		this.counter++
		this.conns[id] = conn

		go func() {
			buf := make([]byte, 1024)
			for {
				msg := &net.TunnelMessageReq{TunID: this.tunID, ConnID: id}

				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("client.Read", err)
					msg.Close = true
					_, _ = this.client.rpcClient.Call(this.client, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout)
					break
				}

				fmt.Println("client.Read", buf[:n])
				msg.Data = buf[:n]
				if _, err := this.client.rpcClient.Call(this.client, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout); err != nil {
					break
				}

			}

			this.client.taskQueue.Push(func() {
				this.close(id)
			})
		}()
	})
}
*/
