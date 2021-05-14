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

	counter   uint32
	slaves    map[uint32]*slave
	clients   map[uint32]*client
	mapping   map[uint32]*net.Mapping
	listeners map[uint32]*listener
	channels  map[uint32]*channel
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
		slaves:    map[uint32]*slave{},
		clients:   map[uint32]*client{},
		mapping:   map[uint32]*net.Mapping{},
		listeners: map[uint32]*listener{},
		channels:  map[uint32]*channel{},
	}

	this.rpcServer.Register(proto.MessageName(&net.LoginReq{}), this.onLogin)
	this.rpcServer.Register(proto.MessageName(&net.AuthReq{}), this.onAuth)
	this.rpcServer.Register(proto.MessageName(&net.RegisterReq{}), this.onRegister)
	this.rpcServer.Register(proto.MessageName(&net.UnregisterReq{}), this.onUnregister)
	this.rpcServer.Register(proto.MessageName(&net.CloseChannelReq{}), this.onCloseChannel)
	this.rpcServer.Register(proto.MessageName(&net.ChannelMessageReq{}), this.onChannelMessage)
	//this.rpcServer.Register(proto.MessageName(&net.ClientCmdReq{}), this.onClientCommand)

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
							err = this.rpcServer.OnRPCRequest(&endpoint{session: session}, data.(*drpc.Request))
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
		switch ctx.(type) {
		case *slave:
			s := ctx.(*slave)
			remMapIDs := []uint32{}
			for mapId, m := range this.mapping {
				if m.GetSlaveId() == s.id {
					remMapIDs = append(remMapIDs, mapId)
				}
			}

			channels := []uint32{}
			for _, mapId := range remMapIDs {
				delete(this.mapping, mapId)
				if l, ok := this.listeners[mapId]; ok {
					l.destroy()
					delete(this.listeners, mapId)
				}
				for cId, c := range this.channels {
					if c.mapID == mapId {
						channels = append(channels, cId)
					}
				}
			}

			for _, cId := range channels {
				if c, ok := this.channels[cId]; ok {
					c.close()
					delete(this.channels, cId)
				}
			}

			s.session.SetContext(nil)
			s.session = nil
			delete(this.slaves, s.id)
		case *client:
			c := ctx.(*client)
			c.session.SetContext(nil)
			c.session = nil
			delete(this.clients, c.id)
		}
	}
}

func (this *IncMaster) Stop() {

}

func (this *IncMaster) onLogin(replier *drpc.Replier, req interface{}) {
	end := replier.Channel.(*endpoint)
	msg := req.(*net.LoginReq)

	id := this.counter
	this.counter++

	s := &slave{
		endpoint: &endpoint{
			id:      id,
			session: end.session,
		},
		name: msg.GetName(),
	}

	this.slaves[id] = s
	s.session.SetContext(s)
	fmt.Println("onLogin slave", id)
	_ = replier.Reply(&net.LoginResp{Id: id}, nil)
}

func (this *IncMaster) onAuth(replier *drpc.Replier, req interface{}) {
	end := replier.Channel.(*endpoint)
	msg := req.(*net.AuthReq)

	if this.token != "" && this.token != msg.GetToken() {
		replier.Reply(&net.AuthResp{Msg: "token failed"}, nil)
		return
	}

	id := this.counter
	this.counter++

	c := &client{&endpoint{
		id:      id,
		session: end.session,
	}}
	this.clients[id] = c
	c.session.SetContext(c)
	fmt.Println("onAuth client", id)
	_ = replier.Reply(&net.AuthResp{Id: id}, nil)
}

func (this *IncMaster) onRegister(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.RegisterReq)
	fmt.Println("onRegister", msg)

	s, ok := this.slaves[msg.GetSlaveId()]
	if !ok {
		replier.Reply(&net.RegisterResp{Msg: "slave is not exist "}, nil)
		return
	}

	mapId := this.counter
	this.counter++

	// listener
	l := &listener{mapID: mapId}
	address := fmt.Sprintf("%s:%d", this.ip, msg.GetMaps().GetExternalPort())
	if err := l.listen(address, func(conn net2.Conn) {
		this.taskQueue.Push(func() {
			open := &net.OpenChannelReq{MapId: mapId}
			open.ChannelId = this.counter
			this.counter++

			this.rpcClient.Go(s, proto.MessageName(open), open, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
				if e != nil {
					return
				}
				resp := i.(*net.OpenChannelResp)
				if resp.GetMsg() != "" {
					fmt.Println("openChannel error", resp.GetMsg())
					_ = conn.Close()
					return
				}

				ch := &channel{
					channelID: open.GetChannelId(),
					mapID:     open.GetMapId(),
					conn:      conn,
					rpcClient: this.rpcClient,
					session:   s.session,
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

	dialerAddr := fmt.Sprintf("%s:%d", msg.GetMaps().GetInternalIp(), msg.GetMaps().GetInternalPort())
	create := &net.CreateDialerReq{
		MapId:   mapId,
		Address: dialerAddr,
	}

	this.rpcClient.Go(s, proto.MessageName(create), create, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			l.destroy()
			replier.Reply(&net.RegisterResp{Msg: e.Error()}, nil)
			return
		}

		resp := i.(*net.CreateDialerResp)
		if resp.GetMsg() != "" {
			l.destroy()
			replier.Reply(&net.RegisterResp{Msg: resp.GetMsg()}, nil)
			return
		}

		mapInfo := msg.GetMaps()
		mapInfo.MapId = mapId
		mapInfo.SlaveId = msg.GetSlaveId()
		this.mapping[mapId] = mapInfo
		this.listeners[mapId] = l
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
	slaveId := m.GetSlaveId()

	s, ok := this.slaves[slaveId]
	if !ok {
		replier.Reply(&net.UnregisterResp{}, nil)
		return
	}

	destroy := &net.DestroyDialerReq{MapId: mapId}
	this.rpcClient.Go(s, proto.MessageName(destroy), destroy, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			replier.Reply(&net.UnregisterResp{Msg: e.Error()}, nil)
			return
		}

		resp := i.(*net.DestroyDialerResp)
		if resp.GetMsg() != "" {
			replier.Reply(&net.UnregisterResp{Msg: e.Error()}, nil)
			return
		}

		delete(this.mapping, mapId)
		if l := this.listeners[mapId]; l != nil {
			l.destroy()
			delete(this.listeners, mapId)
		}
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
	//fmt.Println("onChannelMessage", msg.GetChannelId())

	ch, ok := this.channels[msg.GetChannelId()]
	if !ok {
		replier.Reply(&net.ChannelMessageResp{Msg: "channel is not exit"}, nil)
		return
	}

	if ch.conn != nil {
		if err := ch.writeTo(msg.GetData()); err != nil {
			ch.close()
			delete(this.channels, ch.channelID)
			replier.Reply(&net.ChannelMessageResp{Msg: err.Error()}, nil)
			return
		}
	} else {

	}
	replier.Reply(&net.ChannelMessageResp{}, nil)
}

/*
func (this *IncMaster) onClientCommand(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.ClientCmdReq)

	var b []byte
	switch msg.GetCmd() {
	case "ml":
		ms := make([]*net.Mapping, 0, len(this.mapping))
		for _, v := range this.mapping {
			ms = append(ms, v.info)
			fmt.Println(v.info)
		}
		b, _ = json.Marshal(ms)
	default:

	}
	replier.Reply(&net.ClientCmdResp{Data: b}, nil)
}
*/
