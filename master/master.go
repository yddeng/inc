package master

import (
	"fmt"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dutil/task"
	"github.com/yddeng/intun/net"
	"github.com/yddeng/intun/util"
	"time"
)

type inMaster struct {
	password   string
	acceptor   *dnet.TCPAcceptor
	counter    uint32
	slaves     map[uint32]*inSlave
	clients    map[uint32]*inClient
	taskQueue  *task.TaskQueue
	dispatcher *util.Dispatcher
	//server     *drpc.Server
}

func (this *inMaster) OnConnection(conn dnet.NetConn) {
	fmt.Println("new client", conn.RemoteAddr().String())
	_ = dnet.NewTCPSession(conn,
		dnet.WithTimeout(time.Second*5, 0), // 超时
		dnet.WithCodec(net.NewCodec()),
		dnet.WithErrorCallback(func(session dnet.Session, err error) {
			fmt.Println("onError", err)
		}),
		dnet.WithMessageCallback(func(session dnet.Session, data interface{}) {
			this.taskQueue.Push(func() {
				this.dispatcher.Dispatch(session, data.(*net.Message))
			})
		}),
		dnet.WithCloseCallback(func(session dnet.Session, reason error) {
			fmt.Println("onClose", reason)
		}))
}

type inSlave struct {
	id      uint32
	name    string
	session dnet.Session
}

type inClient struct {
	id      uint32
	session dnet.Session
}

type tunnel struct {
	c2s map[uint32]uint32
	s2c map[uint32]uint32
}

func (this *tunnel) addTunnel(c, s uint32) {
	this.c2s[c] = s
	this.s2c[s] = c
}

func (this *tunnel) remTunnel(c uint32) {
	if s, ok := this.c2s[c]; ok {
		delete(this.s2c, s)
		delete(this.c2s, c)
	}
}

func LaunchMaster(host string, port int) {
	address := fmt.Sprintf(":%d", port)

	m := &inMaster{
		acceptor:   dnet.NewTCPAcceptor(address),
		counter:    1,
		slaves:     map[uint32]*inSlave{},
		clients:    map[uint32]*inClient{},
		dispatcher: util.NewDispatcher(),
		taskQueue:  task.NewTaskQueue(1024),
		//server:     drpc.NewServer(),
	}

	//m.server.Register()

	m.taskQueue.Run()
	m.acceptor.Serve(m)
}

func (this *inMaster) onRegister(session dnet.Session, msg *net.Message) {
	id := this.counter
	this.counter++

	req := msg.Data.(*net.RegisterReq)
	s := &inSlave{id: id, name: req.GetName(), session: session}
	this.slaves[id] = s
	session.SetContext(s)
	session.Send(&net.Message{Data: &net.RegisterResp{ID: id}})
}

func (this *inMaster) onConnect(session dnet.Session, msg *net.Message) {
	req := msg.Data.(*net.ConnectionReq)
	if req.GetPassword() != this.password {
		session.Send(&net.Message{Data: &net.ConnectionResp{
			Msg: "password is failed",
		}})
		return
	}

	id := this.counter
	this.counter++

	c := &inClient{session: session, id: id}
	this.clients[id] = c
	session.SetContext(c)

	session.Send(&net.Message{Data: &net.ConnectionResp{ID: id}})

}
