package main


import (
	"sync"
	"time"
	"fmt"
)

func main() {
/*
	ticker := time.NewTicker()
	ticker.
*/
	var intChan chan int = make(chan int, 10)
	var strChan  chan string = make(chan string, 10)

	var wg sync.WaitGroup
	wg.Add(2)
	//插入数据
	go func() {
		var count int
		for count < 1000 {
			count++
			select {
			case intChan <- 10:
				fmt.Printf("write to int chan succ\n")
			case strChan <- "hello":
				fmt.Printf("write to str chan succ\n")
			default:
				fmt.Printf("all chan is full\n")
				time.Sleep(time.Second)
			}
		}
		wg.Done()
	}()

	//读取数据
	go func() {
		var count int
		for count < 10000 {
			count++
			select {
			case a := <- intChan:
				fmt.Printf("read from int chan succ, a:%d\n", a)
			case <- strChan:
				fmt.Printf("read from str chan succ\n")
			default:
				fmt.Printf("all chan is empty\n")
				time.Sleep(time.Second)
			}
		}

		wg.Done()
	}()

	wg.Wait()

}