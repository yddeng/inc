package inc

import (
	"net"
)

type listener struct {
	mapID    uint32
	listener net.Listener
}

func (this *listener) destroy() {
	_ = this.listener.Close()
}

func (this *listener) listen(addr string, newConn func(conn net.Conn)) (err error) {
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
					return
				}
			}

			newConn(conn)
		}
	}()
	return
}

/*
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
*/
