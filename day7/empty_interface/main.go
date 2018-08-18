package main


import (
	"fmt"
)

func main(){
	var a interface{}
	var b int

	a = b
	fmt.Printf("a=%v a:%T\n", a, a)
	var c float32

	a = c
	fmt.Printf("a=%v a:%T\n", a, a)
}