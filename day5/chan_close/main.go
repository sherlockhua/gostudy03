package main


import (
	"time"
	"fmt"
)

func main(){
	var intChan chan int
	intChan = make(chan int, 10)
	for i := 0; i <10; i++{
		intChan <- i
	}

	close(intChan)
	time.Sleep(time.Second*10)
	for i := 0; i <10; i++{
		//var a int
		//a = <- intChan
		<- intChan
		//fmt.Printf("a=%d\n", a)
	}
}