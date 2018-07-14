package main

import (
	"fmt"
)

func Adder() func(int)int {
	var x int
	return func(d int) int {
		x = x + d
		return x
	}
}

func main() {
	f := Adder()
	fmt.Println(f(100))
	fmt.Println(f(200))

	f1 := Adder()
	fmt.Println(f1(1))
	fmt.Println(f1(2))
}