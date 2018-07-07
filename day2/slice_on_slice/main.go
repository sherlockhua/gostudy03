package main


import (
	"fmt"
)
func main() {
	var a []int
	var b [6]int = [6]int{1,2,3,4,5,6}
	//包括第0个元素、第1个元素
	a = b[0:2]
	fmt.Printf("len(a):%d cap:%d\n", len(a), cap(a))
	a = b[3:6]
	fmt.Printf("len(a):%d cap:%d\n", len(a), cap(a))

	c := a[1:2]
	fmt.Printf("c=%v\n", c)
}