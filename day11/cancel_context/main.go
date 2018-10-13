package main

import (
	"sync"
	"time"
	"fmt"
	"context"
)
var wg sync.WaitGroup


func worker2(ctx context.Context) {
	LOOP:
	for {
		fmt.Printf("worker2\n")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			break LOOP
		default:

		}
	}
}

func worker(ctx context.Context) {
	//ctx2, cancel := context.WithCancel(ctx)
	go worker2(ctx)
	//cancel()

	LOOP:
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			break LOOP
		default:

		}
	}

	
	wg.Done()
}

func main()  {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second*3)
	cancel()
	wg.Wait()
}