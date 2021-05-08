package client

import (
	"fmt"
	"github.com/ozgio/strutil"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dutil/task"
	"github.com/yddeng/intun/net"
	"github.com/yddeng/intun/util"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

var (
	address string
)

var (
	taskQueue  *task.TaskQueue
	dispatcher *util.Dispatcher
	buffer     []byte
	session    dnet.Session
	id         uint32
	slave      uint32
	name       string
)

func Launch(host string, port int) {
	taskQueue = task.NewTaskQueue(128)
	taskQueue.Run()
	buffer = make([]byte, 128)
	dispatcher = util.NewDispatcher()
	dispatcher.RegisterCallBack(22, onConnectionResp)

	address = fmt.Sprintf("%s:%d", host, port)

	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := dnet.DialTCP(addr, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}

	session = dnet.NewTCPSession(conn,
		dnet.WithTimeout(time.Second*5, 0), // 超时
		dnet.WithCodec(net.NewCodec()),
		dnet.WithErrorCallback(func(session dnet.Session, err error) {
			fmt.Println("onError", err)
		}),
		dnet.WithMessageCallback(func(session dnet.Session, data interface{}) {
			taskQueue.Push(func() {
				dispatcher.Dispatch(session, data.(*net.Message))
			})
		}),
		dnet.WithCloseCallback(func(session dnet.Session, reason error) {
			fmt.Println("onClose", reason)
		}))

	fmt.Println("password")
	p := readLine()

	session.Send(&net.Message{
		Data: &net.ConnectionReq{
			Password: p,
		},
	})

	loop1()
}

func onConnectionResp(session dnet.Session, msg *net.Message) {
	resp := msg.Data.(*net.ConnectionResp)

	fmt.Println(resp.GetMsg(), resp.GetID())
	if resp.GetMsg() != "" {
		return
	}

	id = resp.GetID()
	loop1()
}

func readLine() string {
	n, _ := os.Stdin.Read(buffer)
	return string(buffer[:n-1])
}

func readWords() (string, []string, int) {
	cmd := readLine()
	words := strutil.Words(cmd)
	wordsLen := len(words)
	return cmd, words, wordsLen
}

//
// loop1
//   |____loop2
//         |____loop3
//

func loop1() {
	fmt.Print("intun>")

	_, words, wordsLen := readWords()

	switch wordsLen {
	case 0:
	case 1:
		switch words[0] {
		case "list":
			loop1()
		case "quit":
			os.Exit(0)
		case "help":
			loop1()
		default:
			loop1()
		}
	case 2:
		switch words[0] {
		case "select":
			num, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Println(err)
				loop1()
				break
			}
			slave = uint32(num)
			loop2()
		default:
			loop1()
		}
	default:
		loop1()
	}
}

func loop2() {
	fmt.Print(fmt.Sprintf("intun[%d]>", slave))

	_, words, wordsLen := readWords()
	switch wordsLen {
	case 0:
		loop2()
	case 1:
		switch words[0] {
		case "quit":
			slave = 0
			loop1()
		case "help":
			loop2()
		default:
			createTunnel(words)
		}
	default:
		createTunnel(words)
	}

}

func createTunnel(words []string) {
	if err := replaceExec(words[0], words); err != nil {
		fmt.Println(err)
	}
	loop2()
}

func replaceExec(name string, argv []string) error {
	binary, err := exec.LookPath(name)
	if err != nil {
		return err
	}

	err = syscall.Exec(binary, argv, []string{})
	if err != nil {
		return err
	}
	return nil
}
