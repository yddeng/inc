package net

import (
	"github.com/yddeng/dutil/protocol"
	"github.com/yddeng/dutil/protocol/protobuf"
)

var pb *protocol.Protocol

const (
	CmdRegister   = 1
	CmdConnection = 2
	CmdCommand    = 3
	CmdTunMsg     = 4
	CmdHeart
)

func init() {
	pb = protocol.NewProtoc(&protobuf.Protobuf{})

	pb.Register(11, &RegisterReq{})
	pb.Register(12, &RegisterResp{})

	pb.Register(21, &ConnectionReq{})
	pb.Register(22, &ConnectionResp{})

	pb.Register(31, &CommandReq{})
	pb.Register(32, &CommandResp{})

	pb.Register(40, &TunMsg{})

	pb.Register(50, &Heartbeat{})

}
