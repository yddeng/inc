package inc

import (
	"fmt"
	"github.com/yddeng/utils/task"
	"net"
)

type listener struct {
	mapID     uint32
	taskQueue *task.TaskQueue
	listener  net.Listener
	counter   uint32
	conns     map[uint32]*tcpConn
}

type tcpConn struct {
	connID uint32
	net.Conn
}

func (this *listener) destroy() {
	this.listener.Close()
}

func (this *listener) closeConn(connID uint32) {
	c, ok := this.conns[connID]
	if ok {
		fmt.Printf("conn %d disconnect. ", connID)
		_ = c.Close()
		delete(this.conns, connID)
	}
}

func (this *listener) listen(addr string, newConn func(conn *tcpConn)) (err error) {
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
			this.taskQueue.Push(func() {
				id := this.counter
				this.counter++
				tc := &tcpConn{connID: id, Conn: conn}
				this.conns[id] = tc
				go newConn(tc)
			})
		}
	}()
	return
}

func (this *listener) handleConn(conn *tcpConn, f func(b []byte, close bool)) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read", conn.connID, err)
			f(nil, true)
			break
		}

		fmt.Println("conn.Read", conn.connID, buf[:n])
		f(buf[:n], false)
	}

	this.taskQueue.Push(func() {
		this.closeConn(conn.connID)
	})
}

func (this *listener) writeTo(id uint32, b []byte) error {
	c, ok := this.conns[id]
	if !ok {
		return fmt.Errorf("conn %d is not found", id)
	}

	_, err := c.Write(b)
	return err
}
