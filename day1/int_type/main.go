package main

import (
	"fmt"
)


func main() {
	//var a int = 10
	//var a = 10
	a := 10
	var b int32 = 10
	var c int 
	c = a + int(b)

	//int(b)把b强制转换成int类型
	var d int = int(b)
	fmt.Print(a, b, c, d)
	fmt.Println()
	fmt.Printf("%x\n", c)
}
