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
