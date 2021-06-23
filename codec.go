package inc

import (
	"encoding/binary"
	"github.com/yddeng/utils/protocol"
	"github.com/yddeng/utils/protocol/protobuf"
	"io"
)

const (
	cmdSize  = 1
	seqSize  = 8
	bodySize = 2
	HeadSize = cmdSize + seqSize + bodySize
)

type header [HeadSize]byte

func (hdr header) Cmd() uint8 {
	return hdr[0]
}

func (hdr header) Seqno() uint64 {
	return binary.BigEndian.Uint64(hdr[1:])
}

func (hdr header) Length() uint16 {
	return binary.BigEndian.Uint16(hdr[9:])
}

type Message struct {
	Cmd  uint8
	Seq  uint64
	Data interface{}
}

func ReadMessage(reader io.Reader) (*Message, error) {
	var hdr header
	if _, err := io.ReadFull(reader, hdr[:]); err != nil {
		return nil, err
	}

	data := make([]byte, hdr.Length())
	if _, err := io.ReadFull(reader, data); err != nil {
		return nil, err
	}

	pMsg, err := pbProto.Unmarshal(uint16(hdr.Cmd()), data)
	if err != nil {
		return nil, err
	}

	return &Message{
		Cmd:  hdr.Cmd(),
		Seq:  hdr.Seqno(),
		Data: pMsg,
	}, nil
}

func WriteMessage(writer io.Writer, msg *Message) (int, error) {
	cmd, data, err := pbProto.Marshal(msg.Data)
	if err != nil {
		return 0, err
	}

	buf := make([]byte, HeadSize+len(data))
	buf[0] = uint8(cmd)
	binary.BigEndian.PutUint64(buf[1:], msg.Seq)
	binary.BigEndian.PutUint16(buf[9:], uint16(len(data)))
	copy(buf[HeadSize:], data)

	return writer.Write(buf)
}

func (this *RpcResp) OK() bool {
	return this.GetMsg() == ""
}

var (
	pbProto *protocol.Protocol
)

const (
	CmdLogin      = 101
	CmdAuth       = 102
	CmdRegister   = 103
	CmdUnregister = 104
	CmdDial       = 105
	CmdStreamData = 106

	CmdRPCResp = 200
)

func init() {
	pbProto = protocol.NewProtoc(&protobuf.Protobuf{})

	pbProto.Register(CmdLogin, &LoginReq{})
	pbProto.Register(CmdAuth, &AuthReq{})
	pbProto.Register(CmdRegister, &RegisterReq{})
	pbProto.Register(CmdUnregister, &UnregisterReq{})
	pbProto.Register(CmdDial, &DialReq{})
	pbProto.Register(CmdStreamData, &StreamData{})

	pbProto.Register(CmdRPCResp, &RpcResp{})

}
