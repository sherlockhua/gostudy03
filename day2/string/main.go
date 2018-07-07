package main

import (
	"fmt"
)

func main() {
	var str string
	str = "abc汉子"
	//str[0] = 'c'
	//把str强制转成[]byte数组
	var b []byte = []byte(str)
	//把str强制转成[]utf-8数组
	var chars []rune = []rune(str)

	fmt.Printf("b = %v, len(str)=%d\n", b, len(str))
	fmt.Printf("%c\n", 97)
	fmt.Printf("chars count:%d\n", len(chars))
}