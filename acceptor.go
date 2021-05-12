package inc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/utils/task"
	"net"
)

type acceptor struct {
	taskQueue *task.TaskQueue
	id        uint32
	listener  net.Listener
	counter   uint32
	conns     map[uint32]net.Conn
}

func (this *acceptor) close() {
	this.listener.Close()
}

func (this *acceptor) disconnect(id uint32) {
	c, ok := this.conns[id]
	if ok {
		fmt.Printf("conn %d disconnect. ", id)
		_ = c.Close()
		delete(this.conns, id)
	}
}

func (this *acceptor) listen(addr string) (err error) {
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
			this.handleConn(conn)
		}
	}()
	return
}

func (this *acceptor) handleConn(conn net.Conn) {
	this.taskQueue.Push(func() {
		id := this.counter
		this.counter++
		this.conns[id] = conn

		go func() {
			buf := make([]byte, 1024)
			for {
				msg := &net.TunnelMessageReq{TunID: this.tunID, ConnID: id}

				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("client.Read", err)
					msg.Close = true
					_, _ = this.client.rpcClient.Call(this.client, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout)
					break
				}

				fmt.Println("client.Read", buf[:n])
				msg.Data = buf[:n]
				if _, err := this.client.rpcClient.Call(this.client, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout); err != nil {
					break
				}

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
