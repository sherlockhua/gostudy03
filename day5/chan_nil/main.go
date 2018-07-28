package main


import (
	"fmt"
	"time"
)

func main() {
	var intChan chan int = make(chan int, 1)
	fmt.Printf("%p\n", intChan)
	go func () {
		//intChan <- 100
		fmt.Printf("insert item end\n")
	}()
	go func () {
		fmt.Printf("start")
		time.Sleep(time.Second*3)
		var a int
		a = <- intChan
		fmt.Printf("a=%d\n", a)
		
	}()

	time.Sleep(time.Second*5)
}