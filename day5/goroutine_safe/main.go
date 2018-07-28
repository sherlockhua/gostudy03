package main


import (
	"time"
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func test1(){
	for i := 0; i < 1000000; i++{
		mutex.Lock()
		count++
		mutex.Unlock()
	}
}

func test2(){
	for i := 0; i < 1000000; i++{
		mutex.Lock()
		count++	
		mutex.Unlock()
	}
}


func main(){
	go test1()
	go test2()

	time.Sleep(time.Second)
	fmt.Printf("count=%d\n", count)
}