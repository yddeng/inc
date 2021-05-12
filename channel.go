package inc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	net2 "github.com/yddeng/inc/net"
	"net"
)

type channel struct {
	channelID uint32
	mapID     uint32
	conn      net.Conn
	rpcClient *drpc.Client
	session   dnet.Session
}

func (this *channel) close() {
	if this.conn != nil {
		fmt.Println("channel", this.channelID, "close")
		this.conn.Close()
		this.conn = nil
	}
}

func (this *channel) handleRead(onClose func()) {
	if this.conn != nil {
		buf := make([]byte, 1024)
		for {
			n, err := this.conn.Read(buf)
			if err != nil {
				//fmt.Println("client.Read", err)
				msg := &net2.CloseChannelReq{ChannelId: this.channelID}
				_, _ = this.rpcClient.Call(&endpoint{session: this.session}, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout)
				break
			}

			//fmt.Println("client.Read", buf[:n])
			msg := &net2.ChannelMessageReq{ChannelId: this.channelID, Data: buf[:n]}
			if _, err := this.rpcClient.Call(&endpoint{session: this.session}, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout); err != nil {
				break
			}
		}

		onClose()
	}
}

func (this *channel) writeTo(b []byte) (err error) {
	if this.conn != nil {
		_, err = this.conn.Write(b)
	}
	return
}
