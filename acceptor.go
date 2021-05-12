package inc

import (
	"fmt"
	"github.com/yddeng/utils/task"
	"net"
)

type acceptor struct {
	channelID uint32
	taskQueue *task.TaskQueue
	listener  net.Listener
	counter   uint32
	conns     map[uint32]*tcpConn
}

type tcpConn struct {
	connID uint32
	net.Conn
}

func (this *acceptor) close() {
	this.listener.Close()
}

func (this *acceptor) disconnect(connID uint32) {
	c, ok := this.conns[connID]
	if ok {
		fmt.Printf("conn %d disconnect. ", connID)
		_ = c.Close()
		delete(this.conns, connID)
	}
}

func (this *acceptor) listen(addr string, f func(b []byte, close bool)) (err error) {
	if this.listener, err = net.Listen("tcp", addr); err != nil {
		return
	}

	go func() {
		for {
			conn, err := this.listener.Accept()
			if err != nil {
				if ne, ok := err.(net.Error); ok && ne.Temporary() {
					continue
				} else {
					panic(err)
				}
			}
			this.handleConn(conn, f)
		}
	}()
	return
}

func (this *acceptor) handleConn(conn net.Conn, f func(b []byte, close bool)) {
	this.taskQueue.Push(func() {
		id := this.counter
		this.counter++
		this.conns[id] = &tcpConn{connID: id, Conn: conn}

		go func() {
			buf := make([]byte, 1024)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("client.Read", err)
					f(nil, true)
					break
				}

				fmt.Println("client.Read", buf[:n])
				f(buf[:n], false)
			}

			this.taskQueue.Push(func() {
				this.disconnect(id)
			})
		}()
	})
}

func (this *acceptor) writeTo(id uint32, b []byte) error {
	c, ok := this.conns[id]
	if !ok {
		return fmt.Errorf("conn %d is not found", id)
	}

	_, err := c.Write(b)
	return err
}
