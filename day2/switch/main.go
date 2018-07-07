package main

import (
	"fmt"
)

func main() {
	f := 1
	switch f {
	case 1:
		fmt.Printf("enter case 1")
		fmt.Printf("f =1\n")
		//fallthrough
	case 2:
		fmt.Printf("enter 2 case\n")
		fmt.Printf("f=2\n")
	}
}