package main


import (
	"fmt"
)

func main() {
	var a []int
	var b [6]int = [6]int{1,2,3,4,5,6}
	//包括第0个元素、第1个元素
	a = b[0:2]
	fmt.Printf("a=%v\n", a)
	
	//包括第0个元素一直到最后一个元素
	a = b[0:]
	fmt.Printf("a=%v\n", a)

	//包括第2个元素一直到最后一个元素
	a = b[2:]
	fmt.Printf("a=%v\n", a)

	//包括第2个元素一直到最后一个元素
	a = b[:2]
	fmt.Printf("a=%v\n", a)

	//包括第0个元素一直到最后一个元素
	a = b[:]
	fmt.Printf("a=%v\n", a)
}