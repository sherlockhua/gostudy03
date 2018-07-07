package main

import (
	"fmt"
)

func main() {
	b := [...]int{1,23,3,4,5,5,66}
	for index, value := range b {
		fmt.Printf("index:%d value:%d\n", index, value)
	}
}