package main

import (
	"fmt"
)

func main() {
	var a []int
	if a == nil {
		fmt.Printf("a is null addr\n")
	}
	//a[0] = 100，会panic，因为a是个空切片
	fmt.Printf("a:%v addr:%p\n", a, a)
	a = append(a, 2,3,3)
	fmt.Printf("a:%v addr:%p\n", a, a)

	var b []int = []int{1,2,3}
	//把切片b展开成一个个元素
	a = append(a, b...)
	fmt.Printf("appand slice:%v\n", a)
}