package main

import (
	"fmt"
	_ "github.com/gostudy03/day1/calc"
)

func Calc(a int, b int)(int, int) {
	return a + b, a -b
}

func main() {
	var c int

	a1, b1 := Calc(2, 5)
	fmt.Printf("a1=%d b1=%d\n", a1, b1)
	//go calc.Add(2, 3)
	fmt.Printf("c=%d\n", c)
}