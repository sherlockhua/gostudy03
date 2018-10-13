package main

import (
	"sync"
	"time"
	"fmt"
)

var wg sync.WaitGroup

var exit bool



func worker(exitChan chan struct{}) {
	LOOP:
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Second)
		/*
		if exit {
			break
		}
		*/
		select {
		case <- exitChan:
			break LOOP
		default:

		}

	}

	wg.Done()
}

func main()  {

	var exitChan chan struct{} = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second*3)
	//exit = true
	exitChan <- struct{}{}
	wg.Wait()
}