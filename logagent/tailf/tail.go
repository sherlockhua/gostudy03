package tailf

import (
	"context"
	"github.com/gostudy03/xlog"
	"github.com/hpcloud/tail"
	"github.com/gostudy03/logagent/kafka"
	"encoding/json"
	"github.com/gostudy03/logagent/common"
	"fmt"
)

type TailTask struct {
	Path string
	ModuleName string
	Topic string
	tailx *tail.Tail
	ctx context.Context
	cancel context.CancelFunc
}

var localIP string

func init() {
	var err error
	localIP, err = common.GetLocalIP()
	if err != nil {
		xlog.LogError("get local ip failed, err:%v", err)
		panic(fmt.Sprintf("get local ip failed, er:%v", err))
	}
}


func NewTailTask(path, module, topic string)(tailTask *TailTask, err error) {
	tailTask = &TailTask{}
	err = tailTask.Init(path, module, topic)
	return
}

func (t *TailTask) Init(path, module, topic string) (err error) {
	
	t.Path = path
	t.ModuleName = module
	t.Topic = topic

	t.ctx, t.cancel = context.WithCancel(context.Background())
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

func (t *TailTask) Key() string{
	key := fmt.Sprintf("%s_%s_%s", t.Path, t.ModuleName, t.Topic)
	return key
}

func (t *TailTask) Stop() {
	t.cancel()
}

func (t *TailTask) Run() {
	for {
		select {
		case <- t.ctx.Done():
			xlog.LogWarn("task path:%s module:%s topic:%s is exit", t.Path, t.ModuleName, t.Topic)
			return
		case line, ok := <- t.tailx.Lines:
			if !ok {
				xlog.LogWarn("get message from tailf failed")
				continue
			}

			if len(line.Text) == 0 {
				continue
			}

			xlog.LogDebug("line:%s", line.Text)
			data := &common.LogAgentData{
				IP: localIP,
				Data: line.Text,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				continue
			}

			msg := &kafka.Message{
				Data: string(jsonData),
				Topic: t.Topic,
			}
	
			err = kafka.SendLog(msg)
			if err != nil {
				xlog.LogWarn("send log failed, err:%v\n", err)
				continue
			}
			xlog.LogDebug("send to kafka succ\n")
		}
	}
}
