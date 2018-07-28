package main

import (
	"fmt"
//	"time"
)

func main() {
	ch := make(chan string)
	exitChan := make (chan bool, 3)
	go sendData(ch, exitChan)
	go getData(ch, exitChan)
	go getData2(ch, exitChan)
	
	//等待其他goroutine退出
	<- exitChan
	<- exitChan
	<- exitChan
	fmt.Printf("main goroutine exited\n")
}

func sendData(ch chan string, exitCh chan bool) {
	ch <- "aaa"
	ch <- "bbb"
	ch <- "ccc"
	ch <- "ddd"
	ch <- "eee"
	close(ch)
	fmt.Printf("send data exited")
	exitCh <- true
}

func getData (ch chan string, exitCh chan bool) {
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
	exitCh <- true
}

func getData2 (ch chan string, exitCh chan bool) {
	//var input2 string
	for {
		//input2 = <- ch
		input2, ok := <- ch
		if !ok {
			break
		}
		// 此处 打印出来的顺序 和写入的顺序 是一致的
		// 遵循队列的原则: 先入先出
		fmt.Printf("getData2中的input值:%s\n", input2)
	}
	fmt.Printf("get data2 exited\n")
	exitCh <- true
}
