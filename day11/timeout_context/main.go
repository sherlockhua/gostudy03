package main

import (
	"sync"
	"time"
	"fmt"
	"context"
)
var wg sync.WaitGroup


func worker(ctx context.Context) {
	
	LOOP:
	for {
		fmt.Printf("worker\n")
		time.Sleep(time.Millisecond*10)
		select {
		case <- ctx.Done():
			break LOOP
		default:

		}
	}

	fmt.Printf("worker done\n")
	wg.Done()
}

func main()  {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second*8)
	cancel()
	wg.Wait()
}