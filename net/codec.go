package net

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/dutil/buffer"
	"io"
)

var ErrTooLarge = fmt.Errorf("Message too large")

const (
	cmdSize  = 2
	ttSize   = 1
	seqSize  = 8
	bodySize = 2
	HeadSize = cmdSize + ttSize + seqSize + bodySize
	BuffSize = 1<<16 - 1
)

const (
	rpc_req  = 0x1
	rpc_resp = 0x2
)

type Codec struct {
	readBuf  *buffer.Buffer
	readHead bool
	cmd      uint16
	tt       byte
	seq      uint64
	body     uint16
}

func NewCodec() *Codec {
	return &Codec{
		readBuf: buffer.NewBufferWithCap(BuffSize),
	}
}

//解码
func (decoder *Codec) Decode(reader io.Reader) (interface{}, error) {
	for {
		msg, err := decoder.unPack()

		if msg != nil {
			return msg, nil

		} else if err == nil {
			_, err1 := decoder.readBuf.ReadFrom(reader)
			if err1 != nil {
				return nil, err1
			}
		} else {
			return nil, err
		}
	}
}

func (decoder *Codec) unPack() (interface{}, error) {

	if !decoder.readHead {
		if decoder.readBuf.Len() < HeadSize {
			return nil, nil
		}

		decoder.cmd, _ = decoder.readBuf.ReadUint16BE()
		decoder.tt, _ = decoder.readBuf.ReadByte()
		decoder.seq, _ = decoder.readBuf.ReadUint64BE()
		decoder.body, _ = decoder.readBuf.ReadUint16BE()
		decoder.readHead = true
	}

	if decoder.body > BuffSize-HeadSize {
		return nil, ErrTooLarge
	}
	if decoder.readBuf.Len() < int(decoder.body) {
		return nil, nil
	}

	data, _ := decoder.readBuf.ReadBytes(int(decoder.body))

	decoder.readHead = false
	switch decoder.tt {
	case rpc_req:
		pMsg, err := pbReq.Unmarshal(decoder.cmd, data)
		if err != nil {
			return nil, err
		}

		return &drpc.Request{
			Seq:    decoder.seq,
			Method: proto.MessageName(pMsg.(proto.Message)),
			Data:   pMsg,
		}, nil

	case rpc_resp:
		pMsg, err := pbResp.Unmarshal(decoder.cmd, data)
		if err != nil {
			return nil, err
		}

		return &drpc.Response{
			Seq:  decoder.seq,
			Data: pMsg,
		}, nil

	default:
		return nil, nil
	}
}

//编码
func (encoder *Codec) Encode(o interface{}) ([]byte, error) {
	var cmd uint16
	var data []byte
	var err error
	var tt byte
	var seq uint64

	switch o.(type) {
	case *drpc.Request:
		tt = rpc_req
		msg := o.(*drpc.Request)
		seq = msg.Seq
		cmd, data, err = pbReq.Marshal(msg.Data)
	case *drpc.Response:
		tt = rpc_resp
		msg := o.(*drpc.Response)
		seq = msg.Seq
		cmd, data, err = pbResp.Marshal(msg.Data)
	}
	if err != nil {
		return nil, err
	}

	bodyLen := len(data)
	if bodyLen > BuffSize-HeadSize {
		return nil, ErrTooLarge
	}

	totalLen := HeadSize + bodyLen
	buff := buffer.NewBufferWithCap(totalLen)
	buff.WriteUint16BE(cmd)
	buff.WriteByte(tt)
	buff.WriteUint64BE(seq)
	buff.WriteUint16BE(uint16(bodyLen))
	buff.WriteBytes(data)

	return buff.Bytes(), nil
}
