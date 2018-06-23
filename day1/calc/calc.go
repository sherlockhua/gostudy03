package calc

import (
	"fmt"
)

func Add(a int, b int) int {
	var c = a +b
	fmt.Printf("add result:%d\n", c)
	return c
}