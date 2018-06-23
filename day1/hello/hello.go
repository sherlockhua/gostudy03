package main

import (
	"fmt"
)

//程序执行的入口函数
func main() {
	fmt.Print("hello world")
	fmt.Printf("hello world, %d\n", 100)
	fmt.Println("hello world")

	var a int
	fmt.Scan(&a)

	test()
}
