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
	CmdLogin = 101
	CmdAuth
	CmdRegister
	CmdUnregister
	CmdCreateDialer
	CmdDestroyDialer
	CmdOpenChannel
	CmdCloseChannel
	CmdChannelMessage
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
