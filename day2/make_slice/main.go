package main

import (
	"fmt"
)

func main() {
	var a []int
	a = make([]int, 2, 2)
	//a[0] = 100
	fmt.Printf("%v len:%d cap:%d, addr:%p\n", a, len(a), cap(a), a)
	a = append(a, 100)
	fmt.Printf("%v len:%d cap:%d, addr:%p\n", a, len(a), cap(a), a)

	/*
	a = append(a, 200, 300, 500, 300)
	fmt.Printf("%v len:%d cap:%d, addr:%p\n", a, len(a), cap(a), a)
	*/
}