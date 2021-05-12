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
