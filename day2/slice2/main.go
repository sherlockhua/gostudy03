package main


import (
	"fmt"
)
func main() {
	var a []int = make([]int, 5, 5)
	b := a
	fmt.Printf("a = %v addr:%p\n", a, a)
	fmt.Printf("b = %v addr:%p\n", b, b)
	
	a = append(a, 10)
	fmt.Printf("a = %v addr:%p\n", a, a)
	fmt.Printf("b = %v addr:%p\n", b, b)
}