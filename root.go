package inc

/*
import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/inc/net"
	"github.com/yddeng/utils/task"
	"reflect"
)

type Root struct {
	taskQueue *task.TaskQueue
	rpcServer *drpc.Server
	rpcClient *drpc.Client

	address  string
	password string
	acceptor dnet.Acceptor

	counter uint32
	ends    map[uint32]*endpoint
	tunnel  map[string]*rootTunnel
}

type rootTunnel struct {
	cli   *endpoint
	leaf  *endpoint
	tunID string
}

func LaunchRoot(address, password string) *Root {
	taskQueue := task.NewTaskQueue(512)
	taskQueue.Run()

	r := &Root{
		address:   address,
		password:  password,
		taskQueue: taskQueue,
		rpcServer: drpc.NewServer(),
		rpcClient: drpc.NewClient(),
		counter:   1,
		ends:      map[uint32]*endpoint{},
		tunnel:    map[string]*rootTunnel{},
	}

	r.rpcServer.Register(proto.MessageName(&net.LeafRegisterReq{}), r.onLeafRegister)
	r.rpcServer.Register(proto.MessageName(&net.CliAuthReq{}), r.onCliAuth)
	r.rpcServer.Register(proto.MessageName(&net.CliCommandReq{}), r.onCliCommand)
	r.rpcServer.Register(proto.MessageName(&net.CreateTunnelReq{}), r.onCreateTunnel)
	r.rpcServer.Register(proto.MessageName(&net.TunnelMessageReq{}), r.onTunnelMessage)

	r.launch()
	return r
}

func (this *Root) launch() {
	this.acceptor = dnet.NewTCPAcceptor(this.address)

	go func() {
		if err := this.acceptor.ServeFunc(func(conn dnet.NetConn) {
			fmt.Println("new client", conn.RemoteAddr().String())
			_ = dnet.NewTCPSession(conn,
				//dnet.WithTimeout(time.Second*10, 0), // 超时
				dnet.WithCodec(net.NewCodec()),
				dnet.WithMessageCallback(func(session dnet.Session, data interface{}) {
					this.taskQueue.Push(func() {
						var err error
						switch data.(type) {
						case *drpc.Request:
							ctx := session.Context()
							if ctx == nil {
								err = this.rpcServer.OnRPCRequest(&endpoint{session: session}, data.(*drpc.Request))
							} else {
								err = this.rpcServer.OnRPCRequest(ctx.(*endpoint), data.(*drpc.Request))
							}
						case *drpc.Response:
							err = this.rpcClient.OnRPCResponse(data.(*drpc.Response))
						default:
							err = fmt.Errorf("invalid type:%s", reflect.TypeOf(data).String())
						}
						if err != nil {
							fmt.Println(err)
						}
					})
				}),
				dnet.WithCloseCallback(func(session dnet.Session, reason error) {
					this.taskQueue.Push(this.onClose, session, reason)
				}))
		}); err != nil {
			panic(err)
		}
	}()
}

func (this *Root) onClose(session dnet.Session, reason error) {
	fmt.Println("onClose", reason)
	ctx := session.Context()
	if ctx != nil {
		end := ctx.(*endpoint)
		end.session.SetContext(nil)
		end.session = nil
		delete(this.ends, end.uId)
	}
}

func (this *Root) Stop() {

}

func (this *Root) onLeafRegister(replier *drpc.Replier, req interface{}) {
	end := replier.Channel.(*endpoint)
	msg := req.(*net.LeafRegisterReq)

	end.uId = this.counter
	this.counter++

	end.name = msg.GetName()
	end.leaf = true

	this.ends[end.uId] = end
	end.session.SetContext(end)
	fmt.Println("onLeafRegister", end.uId)
	replier.Reply(&net.LeafRegisterResp{ID: end.uId}, nil)
}

func (this *Root) onCliAuth(replier *drpc.Replier, req interface{}) {
	end := replier.Channel.(*endpoint)
	msg := req.(*net.CliAuthReq)

	if this.password != "" && this.password != msg.GetPassword() {
		replier.Reply(&net.CliAuthResp{Msg: "password failed"}, nil)
	}

	end.uId = this.counter
	this.counter++

	this.ends[end.uId] = end
	end.session.SetContext(end)
	fmt.Println("onCliAuth", end.uId)
	replier.Reply(&net.CliAuthResp{ID: end.uId}, nil)
}

func (this *Root) onCliCommand(replier *drpc.Replier, req interface{}) {

}

func (this *Root) onCreateTunnel(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.CreateTunnelReq)
	cli := replier.Channel.(*endpoint)

	leaf, ok := this.ends[msg.GetLeafID()]
	if !ok {
		replier.Reply(&net.CreateTunnelResp{Msg: "leaf is not exist"}, nil)
		return
	}

	this.rpcClient.Go(leaf, proto.MessageName(msg), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			replier.Reply(&net.CreateTunnelResp{Msg: e.Error()}, nil)
			return
		}

		resp := i.(*net.CreateTunnelResp)
		if resp.GetMsg() == "" {
			tunID := resp.GetTunnelID()
			tun := &rootTunnel{
				cli:   cli,
				leaf:  leaf,
				tunID: tunID,
			}
			this.tunnel[tunID] = tun
		}
		replier.Reply(resp, nil)

	})
}

func (this *Root) onTunnelMessage(replier *drpc.Replier, req interface{}) {
	end := replier.Channel.(*endpoint)
	msg := req.(*net.TunnelMessageReq)
	fmt.Println("onTunnelMessage", msg)
	tunID := msg.GetTunID()
	tun, ok := this.tunnel[tunID]
	if !ok {
		replier.Reply(&net.TunnelMessageResp{Msg: "tunnel is not exist"}, nil)
		return
	}

	var toEnd *endpoint
	if end == tun.cli {
		toEnd = tun.leaf
	} else {
		toEnd = tun.cli
	}

	this.rpcClient.Go(toEnd, proto.MessageName(msg), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			replier.Reply(&net.CreateTunnelResp{Msg: e.Error()}, nil)
			return
		}

		resp := i.(*net.TunnelMessageResp)
		if resp.GetMsg() != "" {
			delete(this.tunnel, tunID)
		}
		replier.Reply(resp, nil)

	})
}
*/
