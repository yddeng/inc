package inc

import (
	"fmt"
	"github.com/yddeng/utils/task"
	"net"
)

type dialer struct {
	mapID     uint32
	taskQueue *task.TaskQueue
	address   string
	counter   uint32
	conns     map[uint32]*tcpConn
}

func (this *dialer) destroy() {
}

func (this *dialer) closeConn(connID uint32) {
	c, ok := this.conns[connID]
	if ok {
		fmt.Printf("conn %d disconnect. ", connID)
		_ = c.Close()
		delete(this.conns, connID)
	}
}

func (this *dialer) dial() (*tcpConn, error) {
	conn, err := net.Dial("tcp", this.address)
	if err != nil {
		return nil, err
	}

	id := this.counter
	this.counter++
	tc := &tcpConn{connID: id, Conn: conn}
	this.conns[id] = tc
	return tc, nil
}

func (this *dialer) handleConn(conn *tcpConn, f func(b []byte, close bool)) {
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

func (this *dialer) writeTo(id uint32, b []byte) error {
	c, ok := this.conns[id]
	if !ok {
		return fmt.Errorf("conn %d is not found", id)
	}

	_, err := c.Write(b)
	return err
}
