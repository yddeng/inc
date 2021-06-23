package inc

import (
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/smux"
	"net"
	"sync"
)

// inc-ps
// inc-pc

type channel struct {
	inConn   *smux.Stream
	realConn net.Conn

	closeFn   func()
	closeOnce sync.Once
}

func (this *channel) close() {
	this.closeOnce.Do(func() {
		this.closeFn()
		this.inConn.Close()
		this.realConn.Close()
	})
}

func (this *channel) start(closeFunc func()) {
	this.closeFn = closeFunc
	go this.handleRead()
	go this.handleWrite()
}

func (this *channel) handleRead() {
	buf := make([]byte, 1024)
	for {
		n, err := this.realConn.Read(buf)
		if err != nil {
			break
		}
		//fmt.Println("client.Read", buf[:n])
		if _, err = WriteMessage(this.inConn, &Message{Data: &StreamData{Data: buf[:n]}}); err != nil {
			break
		}
	}
	this.close()
}

func (this *channel) handleWrite() {
	for {
		msg, err := ReadMessage(this.inConn)
		if err != nil {
			break
		}
		if _, err := this.realConn.Write(msg.Data.(*StreamData).GetData()); err != nil {
			break
		}
	}
	this.close()
}

func dial(address string, stream *smux.Stream) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	channel := &channel{
		inConn:   stream,
		realConn: conn,
	}

	channel.start(func() {

	})
	return nil
}

type streamChannel struct {
	stream *smux.Stream
}

func (this *streamChannel) SendRequest(req *drpc.Request) error {
	_, err := WriteMessage(this.stream, &Message{Seq: req.Seq, Data: req.Data})
	return err
}

func (this *streamChannel) SendResponse(resp *drpc.Response) error {
	_, err := WriteMessage(this.stream, &Message{Seq: resp.Seq, Data: resp.Data})
	return err
}
