package tailf

import (
	"github.com/gostudy03/xlog"
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(filename string) (err error) {
	
	tailObj, err = tail.TailFile(filename, tail.Config{
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
	/*
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
	*/
}

func ReadLine() (msg *tail.Line, err error) {
	var ok bool
	msg, ok = <- tailObj.Lines
	if !ok {
		err = fmt.Errorf("read line failed")
		return
	}

	return
}
