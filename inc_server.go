package inc

import (
	"fmt"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/smux"
	"github.com/yddeng/utils/task"
)

type ProxyServer struct {
	ip    string
	port  int
	token string

	acceptor dnet.Acceptor

	counter uint32
	clients map[uint32]*endpoint

	taskQueue *task.TaskQueue

	rpcServer *drpc.Server
	rpcClient *drpc.Client
}

func (this *ProxyServer) launch() {
	this.acceptor = dnet.NewTCPAcceptor(fmt.Sprintf("%s:%d", this.ip, this.port))

	go func() {
		if err := this.acceptor.ServeFunc(func(conn dnet.NetConn) {
			fmt.Println("new client", conn.RemoteAddr().String())
			go func() {
				if !smux.IsSmux(conn, connTimeout) {
					conn.Close()
					fmt.Println("remote connection is not smux. ")
					return
				}
				this.taskQueue.Push(this.handleConnection, conn)
			}()

		}); err != nil {
			panic(err)
		}
	}()
}

func (this *ProxyServer) handleConnection(conn dnet.NetConn) {
	smuxSess := smux.SmuxSession(conn)

	client := &endpoint{
		eId:       this.counter,
		smuxSess:  smuxSess,
		streams:   map[uint16]*smux.Stream{},
		rpcServer: this.rpcServer,
		rpcClient: this.rpcClient,
		taskQueue: this.taskQueue,
	}

	this.counter++
	this.clients[client.eId] = client

	go client.listen(func(stream *smux.Stream) {
		client.streams[stream.StreamID()] = stream
		session := client.handleStream(stream)
		session.SetContext(client)
	})

}

func (this *ProxyServer) dispatch() {

}

func (this *ProxyServer) onLogin(replier *drpc.Replier, req interface{}) {
	session := replier.Channel.(*rpcChannel).session
	msg := req.(*LoginReq)

	end := session.Context().(*endpoint)
	end.name = msg.GetName()

	fmt.Println("onLogin slave", end.eId)
	_ = replier.Reply(&RpcResp{}, nil)
}
