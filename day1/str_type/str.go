package main

import (
	"fmt"
)
func test_str1() {
	var b string = "hello\n\n\n"
	var c = "hello"

	fmt.Printf("b=%v and c = %s\n", b, c)
}
func test_str2() {
	var b string = `
	床前明月光\n，
	疑是地上霜。
	举头望明月，
	低头思故乡。
	`

	fmt.Printf("b %s\n", b)
}
func test_char() {
	var c rune
	c = 20320 
	//c = '你' //'你'的utf8编码是20320，所以下面代码输出‘你’
	fmt.Printf("c=%c\n", c)
}
func main() {
	test_str1()
	test_str2()
	test_char()
}