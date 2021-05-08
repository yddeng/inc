package util

import (
	"github.com/yddeng/dnet"
	"github.com/yddeng/intun/net"
)

type handler func(dnet.Session, *net.Message)

type Dispatcher struct {
	handlers map[uint16]handler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: map[uint16]handler{},
	}
}

func (this *Dispatcher) RegisterCallBack(cmd uint16, callback func(session dnet.Session, msg *net.Message)) {
	_, ok := this.handlers[cmd]
	if ok {
		return
	}
	this.handlers[cmd] = callback
}

func (this *Dispatcher) Dispatch(session dnet.Session, msg *net.Message) {
	cmd := msg.Cmd
	handler, ok := this.handlers[cmd]
	if ok {
		handler(session, msg)
	}
}
