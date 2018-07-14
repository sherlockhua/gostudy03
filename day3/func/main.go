package main

import (
	"os"
	"fmt"
	"github.com/gostudy03/day1/day3/test"
)

var s string = "hello"

func Calc(a, b int)(s1 , s2 int) {
	s1 = a +b
	s2 = a - b
	return
}

func testCalc() {
	var sum int
	var sub int
	sum, sub = Calc(2, 5)
	fmt.Printf("sum:%d sub:%d\n", sum, sub)
}

func Add(a ...int) int {
	fmt.Printf("func args count:%d\n", len(a))

	var sum int
	for index, arg := range a {
		fmt.Printf("arg[%d]=%d\n", index, arg)
		sum = sum + arg
	}

	return sum
}

func testAdd() {
	sum := Add()
	fmt.Printf("sum=%d\n", sum)

	sum = Add(1)
	fmt.Printf("sum=%d\n", sum)

	sum = Add(1, 100)
	fmt.Printf("sum=%d\n", sum)
}

func testDefer() {
	defer fmt.Printf("hello world\n")
	defer fmt.Printf("hello world v2\n")

	fmt.Printf("nihao\n")
	fmt.Printf("nihao v2\n")

	file, err := os.Open("C:/Go/robots.txt")
	/*
	defer func(f *os.File) {
		if f != nil {
			f.Close()
		}
	}(file)
	*/
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	//defer file.Close()

	var buf[4096]byte
	n, err := file.Read(buf[:])
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		//file.Close()
		return
	}

	fmt.Printf("read %d byte succ, content:%s\n", n, string(buf[:]))
	//file.Close()
	return
}

func main(){
	//testCalc()
	//testAdd()
	//testDefer()
	Test()
	fmt.Printf("test.s:%s\n", test.Hello)

}