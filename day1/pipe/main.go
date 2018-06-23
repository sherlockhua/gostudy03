package main

import (
	"fmt"
)

func main() {
	p := make(chan int, 3)
	p <- 1
	p <- 2
	p <- 3
	p <- 4

	var b int
	b = <- p
	fmt.Printf("b = %d\n",b)
}