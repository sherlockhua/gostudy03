package main


import (
	"time"
	"fmt"
)

func main() {
	go func(){
		defer func(){
			err := recover()
			if err != nil {
				fmt.Printf("catch panic exception. err:%v\n", err)
			}
		}()

		var p *int
		*p = 1000
		fmt.Printf("hello")
	}()

	var i int
	for {
		fmt.Printf("%d\n", i)
		time.Sleep(time.Second)
	}
}