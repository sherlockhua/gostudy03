package main

import (
	"sync"
	"time"
	"fmt"
	"context"
)
var wg sync.WaitGroup


func worker(ctx context.Context) {
	
	traceCode, ok := ctx.Value("TRACE_CODE").(string)
	if ok {
		fmt.Printf("trace_code:%s\n", traceCode)
	}

	LOOP:
	for {
		
		fmt.Printf("worker,traceCode:%s\n", traceCode)
		time.Sleep(time.Millisecond*10)
		select {
		case <- ctx.Done():
			break LOOP
		default:

		}
	}

	fmt.Printf("worker done, trace_code:%s\n", traceCode)
	wg.Done()
}

func main()  {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	ctx = context.WithValue(ctx, "TRACE_CODE", "3835834873922992")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second*8)
	cancel()
	wg.Wait()
}