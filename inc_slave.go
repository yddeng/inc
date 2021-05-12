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

	dialers map[uint32]*dialer
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
		dialers:   map[uint32]*dialer{},
		rpcServer: drpc.NewServer(),
		rpcClient: drpc.NewClient(),
	}

	this.rpcServer.Register(proto.MessageName(&net.CreateDialerReq{}), leaf.onCreateDialer)
	this.rpcServer.Register(proto.MessageName(&net.DestroyDialerReq{}), leaf.onDestroyDialer)

	this.taskQueue.Push(func() { this.dial() })
	return this

}

func (this *IncSlave) onCreateDialer(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.CreateDialerReq)
	fmt.Println("onCreateDialer", msg)

	dialer := &dialer{mapID: msg.GetMapId(), address: msg.GetAddress()}

	//test
	conn, err := dialer.dial()
	if err != nil {
		replier.Reply(&net.CreateDialerResp{Msg: "dial error" + err.Error()}, nil)
		return
	}

	dialer.closeConn(conn.connID)

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

	dialer.destroy()
	delete(this.dialers, msg.GetMapId())
	replier.Reply(&net.DestroyDialerResp{}, nil)
}
