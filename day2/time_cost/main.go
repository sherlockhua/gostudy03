package main

import (
	"time"
	"fmt"
)

func test() {
	for i := 0; i < 10000000000;i++ {
		_ = i * i *i + i / 10*10*i
	}
}

func main() {
	//var now time.Time
	//now = time.Now()
	//start := now.UnixNano()
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("cost=%d us\n", (end - start)/1000)
}