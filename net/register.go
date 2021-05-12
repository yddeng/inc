package net

import (
	"github.com/yddeng/utils/protocol"
	"github.com/yddeng/utils/protocol/protobuf"
)

var (
	pbReq  *protocol.Protocol
	pbResp *protocol.Protocol
)

const (
	CmdLogin          = 101
	CmdAuth           = 102
	CmdRegister       = 103
	CmdUnregister     = 104
	CmdCreateDialer   = 105
	CmdDestroyDialer  = 106
	CmdOpenChannel    = 107
	CmdCloseChannel   = 108
	CmdChannelMessage = 109
)

func init() {
	pbReq = protocol.NewProtoc(&protobuf.Protobuf{})
	pbResp = protocol.NewProtoc(&protobuf.Protobuf{})

	// proxy
	pbReq.Register(CmdLogin, &LoginReq{})
	pbReq.Register(CmdAuth, &AuthReq{})
	pbReq.Register(CmdRegister, &RegisterReq{})
	pbReq.Register(CmdUnregister, &UnregisterReq{})
	pbReq.Register(CmdCreateDialer, &CreateDialerReq{})
	pbReq.Register(CmdDestroyDialer, &DestroyDialerReq{})
	pbReq.Register(CmdOpenChannel, &OpenChannelReq{})
	pbReq.Register(CmdCloseChannel, &CloseChannelReq{})
	pbReq.Register(CmdChannelMessage, &ChannelMessageReq{})

	pbResp.Register(CmdLogin, &LoginResp{})
	pbResp.Register(CmdAuth, &AuthResp{})
	pbResp.Register(CmdRegister, &RegisterResp{})
	pbResp.Register(CmdUnregister, &UnregisterResp{})
	pbResp.Register(CmdCreateDialer, &CreateDialerResp{})
	pbResp.Register(CmdDestroyDialer, &DestroyDialerResp{})
	pbResp.Register(CmdOpenChannel, &OpenChannelResp{})
	pbResp.Register(CmdCloseChannel, &CloseChannelResp{})
	pbResp.Register(CmdChannelMessage, &ChannelMessageResp{})

}
