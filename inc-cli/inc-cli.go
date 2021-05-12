package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/inc/net"
	"github.com/yddeng/utils/strutil"
	"github.com/yddeng/utils/task"
	"os"
	"reflect"
	"time"
)

var buffer = make([]byte, 128)

func readLine() string {
	n, _ := os.Stdin.Read(buffer)
	return string(buffer[:n-1])
}

func readWords() (string, []string, int) {
	line := readLine()
	words := strutil.Str2Slice(line)
	wordsLen := len(words)
	return line, words, wordsLen
}

type client struct {
	taskQueue *task.TaskQueue
	rpcServer *drpc.Server
	rpcClient *drpc.Client
	id        uint32
	session   dnet.Session
	channelId uint32
}

func (this *client) SendRequest(req *drpc.Request) error {
	return this.session.Send(req)
}

func (this *client) SendResponse(resp *drpc.Response) error {
	return this.session.Send(resp)
}

func (this *client) exit(err error) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}

func newClient(address, token string) *client {
	conn, err := dnet.DialTCP(address, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	this := &client{
		taskQueue: task.NewTaskQueue(128),
		rpcServer: drpc.NewServer(),
		rpcClient: drpc.NewClient(),
	}

	this.taskQueue.Run()
	this.rpcServer.Register(proto.MessageName(&net.CloseChannelReq{}), this.onCloseChannel)
	this.rpcServer.Register(proto.MessageName(&net.ChannelMessageReq{}), this.onChannelMessage)

	this.session = dnet.NewTCPSession(conn,
		dnet.WithCodec(net.NewCodec()),
		dnet.WithCloseCallback(func(session dnet.Session, reason error) { this.exit(reason) }),
		dnet.WithErrorCallback(func(session dnet.Session, err error) { session.Close(err) }),
		dnet.WithMessageCallback(func(session dnet.Session, data interface{}) {
			this.taskQueue.Push(func() {
				var err error
				switch data.(type) {
				case *drpc.Request:
					err = this.rpcServer.OnRPCRequest(this, data.(*drpc.Request))
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
	req := &net.AuthReq{Token: token}
	ret, err := this.rpcClient.Call(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout)
	if err != nil {
		this.exit(err)
	}

	fmt.Println("inc connection ok\nType 'help' for help.\n")
	resp := ret.(*net.AuthResp)
	if resp.GetMsg() != "" {
		this.exit(fmt.Errorf(resp.GetMsg()))
	}
	this.id = resp.GetId()
	return this
}

func (this *client) onChannelMessage(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.ChannelMessageReq)
	//fmt.Println("onChannelMessage", msg.GetChannelId())
	fmt.Println(string(msg.GetData()))
	replier.Reply(&net.ChannelMessageResp{}, nil)
}

func (this *client) onCloseChannel(replier *drpc.Replier, req interface{}) {
	msg := req.(*net.CloseChannelReq)
	fmt.Println("onCloseChannel", msg)
	this.channelId = 0
	replier.Reply(&net.CloseChannelResp{}, nil)
}

func (this *client) loop() {
	fmt.Print("inc-cli=>")
	line, words, length := readWords()
	switch length {
	case 0:
		this.loop()
	case 1:
		switch words[0] {
		case "quit", "exit", "q":
			return
		case "mlist", "ml":
			fmt.Println(this.command("ml"))
			this.loop()
		case "help", "h":
			outputHelp()
			this.loop()
		default:
			fmt.Println(fmt.Sprintf("invalid command : %s\nTry 'help' for help.\n", line))
			this.loop()
		}

	default:
		fmt.Println(fmt.Sprintf("invalid command : %s\nTry 'help' for help.\n", line))
		this.loop()
	}

}

func main() {
	//signal.Ignore(syscall.SIGINT, syscall.SIGTERM)

	commandLine := flag.NewFlagSet("inc", flag.ExitOnError)
	a := commandLine.String("a", "", "--host=HOSTNAME     start server host, required ")
	commandLine.Parse(os.Args[1:])

	if *a == "" {
		return
	}

	fmt.Print("Password to connection master:")
	pw := readLine()

	client := newClient(*a, pw)
	if client != nil {
		client.loop()
	}
}

func outputHelp() {
	s := `commands:
mlist (alias: ml) ---------- Mapping list on external service.
slist (alias: sl) ---------- Slave list on external service.
exit (alias: quit | q) ----- Exit the process.
select (alias: s) ---------- Select internal service and open channel.
register (alias: r) -------- Register mapping from internal service to external service.
`
	fmt.Println(s)
}

func (this *client) register(inIp string, inPort, exPort int32, desc string) string {
	req := &net.RegisterReq{
		Maps: &net.Mapping{
			InternalIp:   inIp,
			InternalPort: inPort,
			ExternalPort: exPort,
			Description:  desc,
		},
		SlaveId: this.id,
	}

	writC := make(chan string, 1)
	_ = this.rpcClient.Go(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			writC <- e.Error()
			return
		}

		msg := i.(*net.RegisterResp)
		if msg.GetMsg() != "" {
			writC <- msg.GetMsg()
			return
		}

		writC <- ""
	})
	return <-writC
}

func (this *client) unregister(mapId uint32) string {
	req := &net.UnregisterReq{MapId: mapId}
	writC := make(chan string, 1)
	_ = this.rpcClient.Go(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			writC <- e.Error()
			return
		}

		msg := i.(*net.UnregisterResp)
		if msg.GetMsg() != "" {
			writC <- msg.GetMsg()
			return
		}

		writC <- ""
	})
	return <-writC
}

func (this *client) command(cmd string) string {
	req := &net.ClientCmdReq{Cmd: cmd}
	writC := make(chan string, 1)
	_ = this.rpcClient.Go(this, proto.MessageName(req), req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			writC <- e.Error()
			return
		}

		msg := i.(*net.ClientCmdResp)
		switch cmd {
		case "ml":
			writC <- outMappingList(msg.GetData())
		default:
			writC <- ""
		}

	})
	return <-writC
}

func outMappingList(b []byte) string {
	var ms []*net.Mapping
	if err := json.Unmarshal(b, &ms); err != nil {
		return err.Error()
	}

	var s = ""
	for _, v := range ms {
		s += fmt.Sprintf("%d %d %s %d %d %s\n", v.GetMapId(), v.GetSlaveId(), v.GetInternalIp(), v.GetInternalPort(), v.GetExternalPort(), v.GetDescription())
	}
	return s
}
