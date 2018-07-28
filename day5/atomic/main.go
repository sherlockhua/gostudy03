package main

import (
	"time"
	"fmt"
	"sync/atomic"
	"sync"
)

var count int32
var mutex sync.Mutex

func test1(){
	for i := 0; i < 1000000; i++{
		//mutex.Lock()
		//count++
		//mutex.Unlock()
		atomic.AddInt32(&count, 1)
	}
}

func test2(){
	for i := 0; i < 1000000; i++{
		//mutex.Lock()
		//count++	
		//mutex.Unlock()
		atomic.AddInt32(&count, 1)
	}
}


func main(){
	go test1()
	go test2()

	time.Sleep(time.Second)
	fmt.Printf("count=%d\n", count)
}