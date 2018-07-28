package main


import (
	"fmt"
)

func main() {
	var intChan <- chan int = make(chan int, 100)
	//intChan <- 100
	var ch  chan<- int = make(chan int, 100)
	//<-ch
}