package main

import (
	"fmt"
)

func testPoint1() {
	var a *int
	var b int = 200
	
	fmt.Printf("value of a :%v\n", a)
	a = &b
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", *a)
}

func testPoint2() {
	var a *int
	var b int = 200
	
	a = &b
	fmt.Printf("value of a %v\n", a)
	fmt.Printf("address of b %v\n", &b)
	fmt.Printf("address of a %v\n", &a)

	var c int = 300
	a = &c
	fmt.Printf("*a = %v\n", *a)
}
 

func testPoint3() {
	var a *int
	var b int = 200
	
	a = &b
	fmt.Printf("value of a %v\n", a)
	fmt.Printf("address of b %v\n", &b)
	fmt.Printf("address of a %v\n", &a)

	*a = 300
	fmt.Printf("value of b %v\n", b)
	fmt.Printf("type of a %T\n", a)
}

func testPoint4() {
	var a *int
	*a = 100
	fmt.Printf("%d\n", *a)
}

func main() {
	//testPoint1()
	//testPoint2()
	//testPoint3()
	testPoint4()
}