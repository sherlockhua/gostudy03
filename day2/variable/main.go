package main

import (
	"fmt"
)

func main() {
	var a int = 100
	var b int 
	b = a
	b = 200
	fmt.Printf("a=%d b = %d, %p %p\n", a, b, &a, &b)
}