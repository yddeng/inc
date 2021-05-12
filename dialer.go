package inc

import "github.com/yddeng/utils/task"

type dialer struct {
	channelID uint32
	taskQueue *task.TaskQueue
	address   string
	counter   uint32
	conns     map[uint32]*tcpConn
}
