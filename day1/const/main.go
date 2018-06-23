package main

import (
	"fmt"
)

const (
	x = 0
	f = 2
	a = iota
	b 
	c = 3
	d
	f1 = iota
)

func main() {
	//var f = 
	fmt.Printf("a=%d b = %d c=%d d=%d f=%v\n", a, b, c, d, f1)
}