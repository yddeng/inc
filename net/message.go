package net

import "github.com/golang/protobuf/proto"

type Message struct {
	Cmd  uint16
	From uint32
	To   uint32
	Data interface{}
}

func NewMessage(msg proto.Message, to uint32) *Message {
	return &Message{
		To:   to,
		Data: msg,
	}
}
