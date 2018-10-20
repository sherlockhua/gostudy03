package tailf

import (
	"github.com/gostudy03/xlog"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/gostudy03/logagent/kafka"
)

type TailObj struct {
	Path string
	ModuleName string
	Topic string
	tailx *tail.Tail
}

func NewTailObj(path, module, topic string)(tailObj *TailObj, err error) {
	tailObj = &TailObj{}
	err = tailObj.Init(path, module, topic)
	return
}

func (t *TailObj) Init(path, module, topic string) (err error) {
	
	t.Path = path
	t.ModuleName = module
	t.Topic = topic

	t.tailx, err = tail.TailFile(path, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		xlog.LogError("tail file err:", err)
		return
	}
	return
}

func (t *TailObj) Run() {
	for {
		select {
		case line, ok := <- t.tailx.Lines:
			if !ok {
				xlog.LogWarn("get message from tailf failed")
				continue
			}

			if len(line.Text) == 0 {
				continue
			}

			xlog.LogDebug("line:%s", line.Text)
			msg := &kafka.Message{
				Line: line.Text,
				Topic: t.Topic,
			}
	
			err := kafka.SendLog(msg)
			if err != nil {
				xlog.LogWarn("send log failed, err:%v\n", err)
				continue
			}
			xlog.LogDebug("send to kafka succ\n")
		}
	}
}

func (t *TailObj)ReadLine() (msg *tail.Line, err error) {
	var ok bool
	msg, ok = <- t.tailx.Lines
	if !ok {
		err = fmt.Errorf("read line failed")
		return
	}

	return
}
