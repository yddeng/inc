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
	CmdProxyRegister   = 101
	CmdProxyUnregister = 102
	CmdProxyList       = 103
	CmdProxyMessage    = 104
)

const (
	CmdLeafRegister  = 1
	CmdCliAuth       = 2
	CmdCliCommand    = 3
	CmdCreateTunnel  = 4
	CmdTunnelMessage = 5
	CmdHeartbeat     = 6
	CmdCloseTunnel   = 7
)

func init() {
	pbReq = protocol.NewProtoc(&protobuf.Protobuf{})
	pbResp = protocol.NewProtoc(&protobuf.Protobuf{})

	// proxy
	pbReq.Register(CmdProxyRegister, &ProxyRegisterReq{})
	pbReq.Register(CmdProxyUnregister, &ProxyUnregisterReq{})
	pbReq.Register(CmdProxyList, &ProxyListReq{})
	pbReq.Register(CmdProxyMessage, &ProxyMessageReq{})

	pbResp.Register(CmdProxyRegister, &ProxyRegisterResp{})
	pbResp.Register(CmdProxyUnregister, &ProxyUnregisterResp{})
	pbResp.Register(CmdProxyList, &ProxyListResp{})
	pbResp.Register(CmdProxyMessage, &ProxyMessageResp{})

	//
	pbReq.Register(CmdLeafRegister, &LeafRegisterReq{})
	pbReq.Register(CmdCliAuth, &CliAuthReq{})
	pbReq.Register(CmdCliCommand, &CliCommandReq{})
	pbReq.Register(CmdCreateTunnel, &CreateTunnelReq{})
	pbReq.Register(CmdTunnelMessage, &TunnelMessageReq{})
	pbReq.Register(CmdHeartbeat, &Heartbeat{})
	pbReq.Register(CmdCloseTunnel, &CloseTunnelReq{})

	pbResp.Register(CmdLeafRegister, &LeafRegisterResp{})
	pbResp.Register(CmdCliAuth, &CliAuthResp{})
	pbResp.Register(CmdCliCommand, &CliCommandResp{})
	pbResp.Register(CmdCreateTunnel, &CreateTunnelResp{})
	pbResp.Register(CmdTunnelMessage, &TunnelMessageResp{})
	pbResp.Register(CmdHeartbeat, &Heartbeat{})
	pbResp.Register(CmdCloseTunnel, &CloseTunnelResp{})

}
