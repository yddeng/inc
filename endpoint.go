package inc

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/inc/net"
	"github.com/yddeng/smux"
	"github.com/yddeng/utils/task"
	"io"
	net2 "net"
	"reflect"
	"time"
)

type endpoint struct {
	eId uint32
	//server   *ProxyServer
	smuxSess *smux.Session
	streams  map[uint16]*smux.Stream

	rpcServer *drpc.Server
	rpcClient *drpc.Client

	taskQueue *task.TaskQueue

	listeners map[uint32]*listener

	name   string
	isAuth bool
}

func (this *endpoint) start() {
	this.listen(func(stream *smux.Stream) {
		this.streams[stream.StreamID()] = stream
		go handleStreamRead(stream, func(message *Message) bool {
			switch message.Cmd {
			case CmdDial:
				this.taskQueue.WaitPush(func() {
					address := message.Data.()
					dial()
				})
				return false
			case CmdRPCResp:
				this.rpcClient.OnRPCResponse(message.Data.(*drpc.Response))
			default:
				this.rpcServer.OnRPCRequest(&streamChannel{stream: stream}, message.Data.(*drpc.Request))
			}
			return true
		})
	})
}

func (this *endpoint) listen(newStream func(stream *smux.Stream)) {
	for {
		stream, err := this.smuxSess.Accept()
		if err != nil {
			return
		}
		this.taskQueue.Push(newStream, stream)
	}
}

func (this *endpoint) onClose(session dnet.Session, reason error) {
}

func (this *endpoint) auth(msg proto.Message) error {
	stream, err := this.smuxSess.Open()
	if err != nil {
		return err
	}

	return this.rpcClient.Go(&streamChannel{stream: stream}, proto.MessageName(msg), msg, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		resp := i.(*RpcResp)
		if !resp.OK() {
			panic(resp.GetMsg())
		}
		this.isAuth = true
		stream.Close()

		go this.start()
	})
}

func handleStreamRead(stream *smux.Stream, callback func(*Message) bool) {
	for {
		msg, err := ReadMessage(stream)
		if err != nil {
			break
		}

		if !callback(msg) {
			break
		}
	}
}

func DialEncode(conn net2.Conn, auth net.Auth, value string) error {
	req := &net.ConnectionReq{
		Id:    auth,
		Value: value,
	}
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	buf := make([]byte, len(data)+2)
	binary.BigEndian.PutUint16(buf, uint16(len(data)))
	copy(buf[2:], data)

	conn.SetWriteDeadline(time.Now().Add(connTimeout))
	defer conn.SetWriteDeadline(time.Time{})
	if _, err = conn.Write(buf); err != nil {
		return err
	}

	//

	conn.SetReadDeadline(time.Now().Add(connTimeout))
	defer conn.SetReadDeadline(time.Time{})

	var b = make([]byte, 2)
	if _, err = io.ReadFull(conn, b); err != nil {
		return err
	}

	length := binary.BigEndian.Uint16(b)

	b = make([]byte, length)
	if _, err = io.ReadFull(conn, b); err != nil {
		return err
	}

	var resp net.ConnectionResp
	if err = proto.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.GetMsg() != "" {
		return errors.New(resp.GetMsg())
	}
	return nil
}

func acceptDecode(conn net2.Conn, token string) (req *net.ConnectionReq, err error) {
	conn.SetReadDeadline(time.Now().Add(connTimeout))
	defer conn.SetReadDeadline(time.Time{})

	var b = make([]byte, 2)
	if _, err = io.ReadFull(conn, b); err != nil {
		return
	}

	length := binary.BigEndian.Uint16(b)

	b = make([]byte, length)
	if _, err = io.ReadFull(conn, b); err != nil {
		return
	}

	if err = proto.Unmarshal(b, req); err != nil {
		return
	}

	if req.GetId() == net.Auth_cline && token != "" && req.GetValue() != token {
		err = errors.New("token is failed. ")
		return
	}

	//

	var data []byte
	if data, err = proto.Marshal(&net.ConnectionReq{}); err != nil {
		return
	}

	buf := make([]byte, len(data)+2)
	binary.BigEndian.PutUint16(buf, uint16(len(data)))
	copy(buf[2:], data)

	conn.SetWriteDeadline(time.Now().Add(connTimeout))
	defer conn.SetWriteDeadline(time.Time{})

	if _, err = conn.Write(buf); err != nil {
		return
	}
	return
}
