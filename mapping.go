package inc

import "github.com/yddeng/inc/net"

type mapping struct {
	mapId    uint32
	slaveId  uint32
	info     *net.Mapping
	listener *listener
}
