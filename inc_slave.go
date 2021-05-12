package inc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/inc/net"
	"github.com/yddeng/utils/task"
	"reflect"
	"time"
)

type IncSlave struct {
	taskQueue *task.TaskQueue
	rpcServer *drpc.Server
	rpcClient *drpc.Client

	id    uint32
	name  string
	rAddr string

	dialing bool
	session dnet.Session

	dialers  map[uint32]*dialer  // mapId for key,
	channels map[uint32]*channel // channelId for key,
}

func (this *IncSlave) SendRequest(req *drpc.Request) error {
	return this.session.Send(req)
}

func (this *IncSlave) SendResponse(resp *drpc.Response) error {
	return this.session.Send(resp)
}

func (this *IncSlave) dial() {
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

func (this *IncSlave) onConnected(conn dnet.NetConn) {
	this.taskQueue.Push(func() {
		this.dialing = false

		this.session = dnet.NewTCPSession(conn,
			dnet.WithCodec(net.NewCodec()),
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
		req := &net.LoginReq{Name: this.name, Id: this.id}
		if err := this.rpcClient.Go(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
			if e != nil {
				fmt.Printf("onConnected loginResp failed, error %s", e.Error())
				panic(e.Error())
			}

			msg := i.(*net.LoginResp)
			if msg.GetMsg() != "" {
				fmt.Printf("onConnected loginResp false, msg %s", msg.GetMsg())
				panic(msg.GetMsg())
			}

			fmt.Println("onConnected login center ok")
			this.id = msg.GetId()

			this.testInit()
			this.testInit2()
		}); err != nil {
			panic(err)
		}
	})
}

func LaunchIncSlave(name, rootAddr string) *IncSlave {
	taskQueue := task.NewTaskQueue(512)
	taskQueue.Run()

	this := &IncSlave{
		name:      name,
		rAddr:     rootAddr,
		taskQueue: taskQueue,
		rpcServer: drpc.NewServer(),
		rpcClient: drpc.NewClient(),
		dialers:   map[uint32]*dialer{},
		channels:  map[uint32]*channel{},
	}

	this.rpcServer.Register(proto.MessageName(&net.CreateDialerReq{}), this.onCreateDialer)
	this.rpcServer.Register(proto.MessageName(&net.DestroyDialerReq{}), this.onDestroyDialer)
	this.rpcServer.Register(proto.MessageName(&net.OpenChannelReq{}), this.onOpenChannel)
	this.rpcServer.Register(proto.MessageName(&net.CloseChannelReq{}), this.onCloseChannel)
	this.rpcServer.Register(proto.MessageName(&net.ChannelMessageReq{}), this.onChannelMessage)

	this.taskQueue.Push(func() { this.dial() })
	return this

}

func (this *IncSlave) testInit() {
	req := &net.RegisterReq{
		Maps: &net.Mapping{
			InternalIp:   "10.128.2.123",
			InternalPort: 22,
			ExternalPort: 2346,
			Description:  "ssh",
		},
		SlaveId: this.id,
	}

	_ = this.rpcClient.Go(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			fmt.Printf("testInit error %s", e.Error())
			return
		}

		msg := i.(*net.RegisterResp)
		if msg.GetMsg() != "" {
			fmt.Printf("testInit msg %s", msg.GetMsg())
			return
		}

		fmt.Println("testInit ok")
	})

}

func (this *IncSlave) testInit2() {
	req := &net.RegisterReq{
		Maps: &net.Mapping{
			InternalIp:   "127.0.0.1",
			InternalPort: 5432,
			ExternalPort: 2345,
			Description:  "psql",
		},
		SlaveId: this.id,
	}

	_ = this.rpcClient.Go(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			fmt.Printf("testInit error %s", e.Error())
			return
		}

		msg := i.(*net.RegisterResp)
		if msg.GetMsg() != "" {
			fmt.Printf("testInit msg %s", msg.GetMsg())
			return
		}

		fmt.Println("testInit ok")
	})

}

func (this *IncSlave) onCreateDialer(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.CreateDialerReq)
	fmt.Println("onCreateDialer", msg)

	dialer := &dialer{mapID: msg.GetMapId(), address: msg.GetAddress()}

	// test
	conn, err := dialer.dial()
	if err != nil {
		replier.Reply(&net.CreateDialerResp{Msg: "dial error" + err.Error()}, nil)
		return
	}

	_ = conn.Close()

	this.dialers[dialer.mapID] = dialer
	replier.Reply(&net.CreateDialerResp{}, nil)
}

func (this *IncSlave) onDestroyDialer(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.DestroyDialerReq)
	fmt.Println("onDestroyDialer", msg)
	dialer, ok := this.dialers[msg.GetMapId()]
	if !ok {
		// 默认已经关闭
		replier.Reply(&net.DestroyDialerResp{}, nil)
		return
	}

	delete(this.dialers, dialer.mapID)
	replier.Reply(&net.DestroyDialerResp{}, nil)
}

func (this *IncSlave) onOpenChannel(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.OpenChannelReq)
	fmt.Println("onOpenChannel", msg)

	ch := &channel{
		channelID: msg.GetChannelId(),
		mapID:     msg.GetMapId(),
		session:   this.session,
		rpcClient: this.rpcClient,
	}

	if ch.mapID != 0 {
		dialer, ok := this.dialers[msg.GetMapId()]
		if !ok {
			replier.Reply(&net.OpenChannelResp{Msg: "dialer is not exist"}, nil)
			return
		}

		conn, err := dialer.dial()
		if err != nil {
			replier.Reply(&net.OpenChannelResp{Msg: "dialer error" + err.Error()}, nil)
			return
		}
		ch.conn = conn

		// read
		go ch.handleRead(func() {
			this.taskQueue.Push(func() {
				ch.close()
				delete(this.channels, ch.channelID)
			})
		})

	}

	this.channels[ch.channelID] = ch
	replier.Reply(&net.OpenChannelResp{}, nil)
}

func (this *IncSlave) onCloseChannel(replier *drpc.Replier, req interface{}) {
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

func (this *IncSlave) onChannelMessage(replier *drpc.Replier, req interface{}) {
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
			replier.Reply(&net.ChannelMessageResp{Msg: "channel writeTo error" + err.Error()}, nil)
			return
		}
	} else {

	}
	replier.Reply(&net.ChannelMessageResp{}, nil)
}
