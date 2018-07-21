package main

import (
	"fmt"
)
func modify (a int) {
	fmt.Printf("2. address of a=%p, value of a:%v\n", &a, a)
	a = 1000
}

func modify2(a *int) {
	fmt.Printf("4. address of a:%v, value of a :%v\n", &a, a)
	*a = 1000
}


func main() {
	var b int = 100
	fmt.Printf("1. address of b=%p, value of b:%v\n", &b, b)
	modify(b)
	

	var p *int = &b
	fmt.Printf("3. address of p:%v, value of p:%v\n", &p, p)

	modify2(p)
	fmt.Printf("b=%d\n", b)
}