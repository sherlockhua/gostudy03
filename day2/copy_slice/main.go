package main

import (
	"fmt"
	//"flag"
)

func main() {
	a := []int{1,2,3}
	b := []int{10, 11}
	copy( b, a)
	fmt.Printf("a:%v b:%v\n", a, b)
}