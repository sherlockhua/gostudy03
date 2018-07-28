package main

import (
	"fmt"
	"sync"
//	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	//exitChan := make (chan bool, 3)
	wg.Add(3)
	go sendData(ch, &wg)
	go getData(ch, &wg)
	go getData2(ch, &wg)

	wg.Wait()
	fmt.Printf("main goroutine exited\n")
}

func sendData(ch chan string, waitGroup *sync.WaitGroup) {
	ch <- "aaa"
	ch <- "bbb"
	ch <- "ccc"
	ch <- "ddd"
	ch <- "eee"
	close(ch)
	fmt.Printf("send data exited")
	waitGroup.Done()
}

func getData (ch chan string, waitGroup *sync.WaitGroup) {
	//var input string
	for {
		//input = <- ch
		input, ok := <- ch
		if !ok {
			break
		}
		// 此处 打印出来的顺序 和写入的顺序 是一致的
		// 遵循队列的原则: 先入先出
		fmt.Printf("getData中的input值:%s\n", input)
	}
	fmt.Printf("get data exited\n")
	waitGroup.Done()
}

func getData2 (ch chan string, waitGroup *sync.WaitGroup) {
	//var input2 string
	for v := range ch {
		fmt.Printf("get data2 %s\n", v)
	}
	
	fmt.Printf("get data2 exited\n")
	waitGroup.Done()
}
