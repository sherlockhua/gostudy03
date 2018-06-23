package main

import (
	"fmt"
)

func main() {
	var b = false
	var (
		c      = true
		d bool = false
	)

	_ = d
	if !b {
		fmt.Printf("b is false\n")
	}

	if !b && c {
		fmt.Printf("result is true\n")
	}
	/*
		if b || c == true {

		}
	*/
	if b || c {
		fmt.Printf("or operation\n")
	}
}
