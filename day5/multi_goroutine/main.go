package main


import (
	"time"
	"fmt"
)

func numbers() {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Millisecond*250)
		fmt.Printf("%d ", i)
		
	}
}

func chars() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(time.Millisecond*400)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers()
	go chars()

	time.Sleep(3*time.Second)
}