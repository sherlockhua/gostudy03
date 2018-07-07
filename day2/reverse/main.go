package main

import (
	"fmt"
)

func main() {
	var str string = "abcdefg"
	bytes := []byte(str)
	//var i int 
	// i = 0
	//var i = 0
	for i := 0; i < len(bytes)/2; i++ {
		//fmt.Printf("%c ", str[i])
		//var tmp = bytes[i]
		//bytes[i] = bytes[len(bytes)-i-1]
		//bytes[len(bytes)-i-1] = tmp

		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}
	str = string(bytes)
	fmt.Printf("reverse string:%s\n", str)
}