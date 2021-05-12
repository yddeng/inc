package inc

import (
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
)

type endpoint struct {
	leaf    bool
	name    string
	uId     uint32
	session dnet.Session
}

func (this *endpoint) SendRequest(req *drpc.Request) error {
	return this.session.Send(req)
}

func (this *endpoint) SendResponse(resp *drpc.Response) error {
	return this.session.Send(resp)
}
