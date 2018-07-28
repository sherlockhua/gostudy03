package main

import (
	"time"
	"fmt"
)

func hello() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello:%d\n", i)
		time.Sleep(time.Millisecond*10)
	}
}

func main() {
	 go hello()
	
	for i := 0; i < 10; i++ {
		fmt.Printf("main:%d\n", i)
		time.Sleep(time.Millisecond*10)
	}
	time.Sleep(time.Second)
}