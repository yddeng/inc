package inc

import (
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/smux"
	"net"
)

type listener struct {
	mapID    uint32
	listener net.Listener
	smuxSess *smux.Session
	channels map[uint16]*channel

	rpcClient *drpc.Client
}

func (this *listener) close() {
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

			go this.onConnection(conn)
		}
	}()
	return
}

func (this *listener) onConnection(conn net.Conn) {
	req := DialReq{Address: ""}
}
