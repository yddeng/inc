package main

import (
	"fmt"
	"github.com/yddeng/dutil/strutil"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var forwardConn net.Conn

func listen(laddr, raddr string) error {
	lAddr, err := net.ResolveTCPAddr("tcp", laddr)
	if err != nil {
		return err
	}

	rAddr, err := net.ResolveTCPAddr("tcp", raddr)
	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", lAddr)
	if err != nil {
		return err
	}

	fmt.Println("listen ok", laddr, raddr)
	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				continue
			} else {
				return err
			}
		}

		fmt.Println("new client", conn.RemoteAddr().String())
		fmt.Println()
		if forwardConn == nil {
			forwardConn = conn
			go forwardData(rAddr)
		} else {
			conn.Close()
			fmt.Println("already conn")
			continue
		}

	}
}

func forwardData(raddr *net.TCPAddr) {
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("forward", conn.LocalAddr().String(), conn.RemoteAddr().String())
	go func() {
		defer fmt.Println("rconn read close")
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("1rconn", conn.LocalAddr().String(), err)
				break
			}
			fmt.Println("rconn read", n, buf[:n])
			if _, err := forwardConn.Write(buf[:n]); err != nil {
				panic(err)
			}
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := forwardConn.Read(buf)
		if err != nil {
			fmt.Println("1forwardConn", err)
			break
		}
		fmt.Println("forward conn read", n, buf[:n])
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("2forwardConn", err)
			break
		}
	}
	fmt.Println("forwardConn close")
	//conn.Close()
	forwardConn.Close()
	forwardConn = nil
}

func readLine() string {
	buffer := make([]byte, 128)
	n, _ := os.Stdin.Read(buffer)
	return string(buffer[:n-1])
}

func readWords() (string, []string, int) {
	cmd := readLine()
	words := strutil.Str2Slice(cmd)
	wordsLen := len(words)
	return cmd, words, wordsLen
}

// 替换当前进程
func replaceExec(name string, argv []string) error {
	binary, err := exec.LookPath(name)
	if err != nil {
		return err
	}

	err = syscall.Exec(binary, argv, os.Environ())
	if err != nil {
		return err
	}
	return nil
}

func forkExec(name string, argv []string) error {
	binary, err := exec.LookPath(name)
	if err != nil {
		return err
	}

	_, err = syscall.ForkExec(binary, argv, &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	})
	if err != nil {
		return err
	}
	return nil
}

func cmdExec(name string, argv []string) *exec.Cmd {
	cmd := exec.Command(name, argv[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd
}

func startServer() {
	laddr := "127.0.0.1:2345"
	raddr := "10.128.2.123:23455"
	if err := listen(laddr, raddr); err != nil {
		panic(err)
	}
}

var pid int

func main() {
	signal.Ignore(syscall.SIGINT, syscall.SIGTERM)

	go startServer()

loop:
	fmt.Printf("==>")
	_, words, length := readWords()
	switch length {
	case 0:
		goto loop
	case 1:
		switch words[0] {
		case "quit":
			return
		default:
			goto loop
		}
	default:
		cmd := cmdExec(words[0], words)
		cmd.Run()
		goto loop
	}
}
