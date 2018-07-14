package main


import (
	"fmt"
)

func testA1() {
	sum := func(a int, b int) int {
		sum := a + b
		return sum
	}(2, 3)
	fmt.Printf("sum = %d\n", sum)
}

func testA2() {
	f1 := func(a int, b int) int {
		sum := a + b
		return sum
	}

	s1 := f1(2, 5)
	s2 := f1(2, 502)
	fmt.Printf("s1=%d, s2=%d\n", s1, s2)
}

func main () {
	//testA1()
	testA2()
}