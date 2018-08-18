package main

import(
	"fmt"
	"sync"
)


func initProgram(threadNum, chanSize int, waitGroup *sync.WaitGroup) (imageChan chan *Task, err error) {

	if chanSize <= 0 || threadNum <= 0 {
		err = fmt.Errorf("invalid parameter")
		return
	}

	imageChan = make(chan *Task, chanSize)
	for i := 0; i <threadNum; i++ {
		waitGroup.Add(1)
		go procImage(imageChan, waitGroup)
	}

	return
}

func procImage(imageChan chan *Task, wg* sync.WaitGroup) {
	for task := range imageChan {
		err := task.Process()
		if err != nil {
			fmt.Printf("process task:%#v failed, err:%v\n", task, err)
			continue
		}

		fmt.Printf("process task:%#v succ\n", task)
	}
	wg.Done()
}