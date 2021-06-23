package inc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/smux"
	"github.com/yddeng/utils/task"
	"reflect"
	"time"
)

type ProxyClient struct {
	id    uint32
	name  string
	rAddr string

	taskQueue *task.TaskQueue

	rpcServer *drpc.Server
	rpcClient *drpc.Client

	dialing bool

	end *endpoint
}

func (this *ProxyClient) dial() {
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

func (this *ProxyClient) onClose(reason error) {
	fmt.Printf("onConnected session closed, reason: %s\n", reason)
	this.session = nil

	for _, c := range this.channels {
		c.close()
	}

	this.channels = map[uint32]*channel{}
	this.dialers = map[uint32]*dialer{}
	this.dial()
}

func (this *ProxyClient) onConnected(conn dnet.NetConn) {
	this.taskQueue.Push(func() {
		this.dialing = false

		this.end = &endpoint{
			smuxSess:  smux.SmuxSession(conn),
			streams:   map[uint16]*smux.Stream{},
			rpcServer: this.rpcServer,
			rpcClient: this.rpcClient,
			taskQueue: this.taskQueue,
		}

		if err := this.end.auth(&LoginReq{Name: this.name}); err != nil {
			panic(err)
		}
	})
}

func LaunchIncSlave(name, rootAddr string, mappings []*net.Mapping) *ProxyClient {
	taskQueue := task.NewTaskQueue(512)
	taskQueue.Run()

	this := &ProxyClient{
		name:      name,
		rAddr:     rootAddr,
		mappings:  mappings,
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

type clientConn struct {
	pc *ProxyClient
}

func (c *clientConn) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (c *clientConn) Write(p []byte) (n int, err error) {
	return 0, nil
}
