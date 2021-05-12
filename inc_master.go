package inc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/inc/net"
	"github.com/yddeng/utils/task"
	net2 "net"

	"reflect"
)

type IncMaster struct {
	taskQueue *task.TaskQueue
	rpcServer *drpc.Server
	rpcClient *drpc.Client

	ip       string
	port     int
	token    string
	acceptor dnet.Acceptor

	counter  uint32
	mapping  map[uint32]*mapping  // mapId
	channels map[uint32]*channel  // atoi
	ends     map[uint32]*endpoint // atoi
}

func LaunchIncMaster(ip string, port int, token string) *IncMaster {
	taskQueue := task.NewTaskQueue(512)
	taskQueue.Run()

	this := &IncMaster{
		ip:        ip,
		port:      port,
		token:     token,
		taskQueue: taskQueue,
		rpcServer: drpc.NewServer(),
		rpcClient: drpc.NewClient(),
		counter:   1,
		mapping:   map[uint32]*mapping{},
		channels:  map[uint32]*channel{},
		ends:      map[uint32]*endpoint{},
	}

	this.rpcServer.Register(proto.MessageName(&net.LoginReq{}), this.onLogin)
	this.rpcServer.Register(proto.MessageName(&net.AuthReq{}), this.onAuth)
	this.rpcServer.Register(proto.MessageName(&net.RegisterReq{}), this.onRegister)
	this.rpcServer.Register(proto.MessageName(&net.UnregisterReq{}), this.onUnregister)
	this.rpcServer.Register(proto.MessageName(&net.CloseChannelReq{}), this.onCloseChannel)
	this.rpcServer.Register(proto.MessageName(&net.ChannelMessageReq{}), this.onChannelMessage)

	this.launch()
	return this
}

func (this *IncMaster) launch() {
	this.acceptor = dnet.NewTCPAcceptor(fmt.Sprintf("%s:%d", this.ip, this.port))

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

func (this *IncMaster) onClose(session dnet.Session, reason error) {
	fmt.Println("onClose", reason)
	ctx := session.Context()
	if ctx != nil {
		end := ctx.(*endpoint)
		end.session.SetContext(nil)
		end.session = nil
		delete(this.ends, end.uId)
	}
}

func (this *IncMaster) Stop() {

}

func (this *IncMaster) onLogin(replier *drpc.Replier, req interface{}) {
	end := replier.Channel.(*endpoint)
	msg := req.(*net.LoginReq)

	end.uId = this.counter
	this.counter++

	end.name = msg.GetName()
	end.leaf = true

	this.ends[end.uId] = end
	end.session.SetContext(end)
	fmt.Println("onLogin slave", end.uId)
	_ = replier.Reply(&net.LoginResp{Id: end.uId}, nil)
}

func (this *IncMaster) onAuth(replier *drpc.Replier, req interface{}) {
	end := replier.Channel.(*endpoint)
	msg := req.(*net.AuthReq)

	if this.token != "" && this.token != msg.GetToken() {
		replier.Reply(&net.AuthResp{Msg: "token failed"}, nil)
		return
	}

	end.uId = this.counter
	this.counter++

	this.ends[end.uId] = end
	end.session.SetContext(end)
	fmt.Println("onAuth client", end.uId)
	_ = replier.Reply(&net.AuthResp{Id: end.uId}, nil)
}

func (this *IncMaster) onRegister(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.RegisterReq)
	mapId := this.counter
	this.counter++

	fmt.Println("onRegister", msg)

	end, ok := this.ends[msg.GetSlaveId()]
	if !ok {
		replier.Reply(&net.RegisterResp{Msg: "slave is not exist "}, nil)
		return
	}

	// listener
	l := &listener{mapID: mapId}
	address := fmt.Sprintf("%s:%d", this.ip, msg.GetMaps().GetRemotePort())
	if err := l.listen(address, func(conn net2.Conn) {
		this.taskQueue.Push(func() {
			open := &net.OpenChannelReq{MapId: mapId}
			open.ChannelId = this.counter
			this.counter++
			open.AcceptorConnId = this.counter
			this.counter++

			this.rpcClient.Go(end, proto.MessageName(open), open, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
				if e != nil {
					return
				}
				resp := i.(*net.OpenChannelResp)
				if resp.GetMsg() != "" {
					return
				}

				ch := &channel{
					channelID:      open.GetChannelId(),
					acceptorConnID: open.GetAcceptorConnId(),
					mapID:          open.GetMapId(),
					conn:           conn,
					dialerConnID:   resp.GetDialerConnId(),
					rpcClient:      this.rpcClient,
					taskQueue:      this.taskQueue,
					session:        end.session,
				}

				this.channels[ch.channelID] = ch

				go ch.handleRead(func() {
					this.taskQueue.Push(func() {
						ch.close()
						delete(this.channels, ch.channelID)
					})
				})
			})

		})
	}); err != nil {
		replier.Reply(&net.RegisterResp{Msg: "listen error " + err.Error()}, nil)
		return
	}

	dialerAddr := fmt.Sprintf("%s:%d", msg.GetMaps().GetLocalIp(), msg.GetMaps().GetLocalPort())
	create := &net.CreateDialerReq{
		MapId:   mapId,
		Address: dialerAddr,
	}

	this.rpcClient.Go(end, proto.MessageName(create), create, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			l.destroy()
			replier.Reply(&net.RegisterResp{Msg: e.Error()}, nil)
			return
		}

		resp := i.(*net.CreateDialerResp)
		if resp.GetMsg() != "" {
			l.destroy()
			replier.Reply(&net.RegisterResp{Msg: e.Error()}, nil)
			return
		}

		this.mapping[mapId] = &mapping{
			mapId:    mapId,
			slaveId:  msg.GetSlaveId(),
			info:     msg.GetMaps(),
			listener: l,
		}
		replier.Reply(&net.RegisterResp{}, nil)
		fmt.Println("onRegister", msg, "ok")
	})
}

func (this *IncMaster) onUnregister(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.UnregisterReq)
	mapId := msg.GetMapId()

	fmt.Println("onUnregister", msg)
	m, ok := this.mapping[mapId]
	if !ok {
		replier.Reply(&net.UnregisterResp{Msg: "mapping is not exist "}, nil)
		return
	}
	slaveId := m.slaveId

	end, ok := this.ends[slaveId]
	if !ok {
		m.listener.destroy()
		delete(this.mapping, slaveId)
		replier.Reply(&net.UnregisterResp{}, nil)
		return
	}

	destroy := &net.DestroyDialerReq{MapId: mapId}
	this.rpcClient.Go(end, proto.MessageName(destroy), destroy, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			replier.Reply(&net.UnregisterResp{Msg: e.Error()}, nil)
			return
		}

		resp := i.(*net.DestroyDialerResp)
		if resp.GetMsg() != "" {
			replier.Reply(&net.UnregisterResp{Msg: e.Error()}, nil)
			return
		}

		m.listener.destroy()
		delete(this.mapping, slaveId)
		replier.Reply(&net.UnregisterResp{}, nil)
		fmt.Println("onUnregister", msg, "ok")
	})

}

func (this *IncMaster) onCloseChannel(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.CloseChannelReq)
	fmt.Println("onCloseChannel", msg)

	ch, ok := this.channels[msg.GetChannelId()]
	if !ok {
		// 默认已经关闭
		replier.Reply(&net.CloseChannelResp{}, nil)
		return
	}

	ch.close()
	delete(this.channels, ch.channelID)
	replier.Reply(&net.CloseChannelResp{}, nil)
}

func (this *IncMaster) onChannelMessage(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.ChannelMessageReq)
	fmt.Println("onChannelMessage", msg.GetChannelId())

	ch, ok := this.channels[msg.GetChannelId()]
	if !ok {
		replier.Reply(&net.ChannelMessageResp{Msg: "channel is not exit"}, nil)
		return
	}

	if ch.conn != nil {
		if err := ch.writeTo(msg.GetData()); err != nil {
			ch.close()
			delete(this.channels, ch.channelID)
			replier.Reply(&net.ChannelMessageResp{Msg: "channel writeTo error" + err.Error()}, nil)
			return
		}
	} else {

	}
	replier.Reply(&net.ChannelMessageResp{}, nil)
}
