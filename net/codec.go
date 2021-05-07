package net

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dutil/buffer"
	"io"
)

var ErrTooLarge = fmt.Errorf("Message too large")

const (
	cmdSize  = 2
	fromSize = 4
	toSize   = 4
	bodySize = 2
	HeadSize = cmdSize + fromSize + toSize + bodySize
	BuffSize = 1<<16 - 1
)

type Codec struct {
	readBuf  *buffer.Buffer
	readHead bool
	cmd      uint16
	from     uint32
	to       uint32
	bodyLen  uint16
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
		decoder.from, _ = decoder.readBuf.ReadUint32BE()
		decoder.to, _ = decoder.readBuf.ReadUint32BE()
		decoder.bodyLen, _ = decoder.readBuf.ReadUint16BE()
		decoder.readHead = true
	}

	if decoder.bodyLen > BuffSize-HeadSize {
		return nil, ErrTooLarge
	}
	if decoder.readBuf.Len() < int(decoder.bodyLen) {
		return nil, nil
	}

	data, _ := decoder.readBuf.ReadBytes(int(decoder.bodyLen))

	pMsg, err := pb.Unmarshal(decoder.cmd, data)
	if err != nil {
		return nil, err
	}

	msg := &Message{
		Cmd:  decoder.cmd,
		From: decoder.from,
		To:   decoder.to,
		Data: pMsg.(proto.Message),
	}

	decoder.readHead = false
	return msg, nil
}

//编码
func (encoder *Codec) Encode(o interface{}) ([]byte, error) {
	msg := o.(*Message)

	var cmd = msg.Cmd
	var data []byte
	var err error
	switch msg.Data.(type) {
	case []byte:
		data = msg.Data.([]byte)
	case proto.Message:
		cmd, data, err = pb.Marshal(msg.Data)
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
	//cmd
	buff.WriteUint16BE(cmd)
	//from
	buff.WriteUint32BE(msg.From)
	//to
	buff.WriteUint32BE(msg.To)
	//bodylen
	buff.WriteUint16BE(uint16(bodyLen))
	//body
	buff.WriteBytes(data)

	return buff.Bytes(), nil
}
