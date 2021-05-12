package inc

import (
	"net"
)

type dialer struct {
	mapID   uint32
	address string
}

func (this *dialer) dial() (net.Conn, error) {
	return net.Dial("tcp", this.address)
}

/*
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
*/
