package intun

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/dutil/task"
	"github.com/yddeng/intun/net"
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

	id     uint32
	leaf   uint32
	tunnel *clientTunnel
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
	if this.leaf == 0 {
		return fmt.Errorf("select leaf before")
	}

	req := &net.CreateTunnelReq{
		LeafID:  this.leaf,
		Address: raddr,
	}

	ret, err := this.rpcClient.Call(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout)
	if err != nil {
		return err
	}

	resp := ret.(*net.CreateTunnelResp)
	if resp.GetMsg() != "" {
		return fmt.Errorf(resp.GetMsg())
	}

	this.tunnel = &clientTunnel{
		tunID:    resp.GetTunnelID(),
		listener: this.listen(laddr),
	}
	fmt.Println("create tunnel ok", resp.GetTunnelID())
	return nil
}

func (this *Client) onTunnelMessage(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.TunnelMessageReq)
	if this.tunnel == nil {
		replier.Reply(&net.TunnelMessageResp{Msg: "tunnel is not exist"}, nil)
		return
	}

	fmt.Println("--", msg.GetData())
	_, err := this.tunnel.conn.Write(msg.GetData())
	if err != nil {
		replier.Reply(&net.TunnelMessageResp{Msg: "conn.Write " + err.Error()}, nil)
		return
	}
	replier.Reply(&net.TunnelMessageResp{}, nil)
}

type clientTunnel struct {
	tunID    string
	listener net2.Listener
	conn     net2.Conn
}

func (this *Client) listen(laddr string) net2.Listener {
	l, err := net2.Listen("tcp", laddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("listen ok", laddr)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				if ne, ok := err.(net2.Error); ok && ne.Temporary() {
					continue
				} else {
					panic(err)
				}
			}

			if this.tunnel.conn == nil {
				this.tunnel.conn = conn
				go func() {
					buf := make([]byte, 1024)
					for {
						n, err := conn.Read(buf)
						if err != nil {
							conn.Close()
							this.tunnel.conn = nil
							fmt.Println("client.Read", err)
							break
						}

						msg := &net.TunnelMessageReq{
							TunID: this.tunnel.tunID,
							Data:  buf[:n],
						}
						fmt.Println("client.Read", buf[:n])
						if _, err := this.rpcClient.Call(this, proto.MessageName(&net.TunnelMessageReq{}), msg, drpc.DefaultRPCTimeout); err != nil {
							break
						}

					}
				}()
			} else {
				conn.Close()
				fmt.Println("already conn")
				continue
			}
		}
	}()
	return nil
}

func forwardData() {

}
