package main


import (
	"fmt"
)

///定义一个Integer 类型，是int64的别名
type Integer int64

func (i *Integer) Print0() {
	fmt.Printf("i=%d\n", i)
}

func (i *Integer) Set(b int64) {
	*i = Integer(b)
}

func main () {
	var a Integer
	a = 1000

	fmt.Printf("a=%v\n", a)

	var b int64 = 500
	a = Integer(b)
	fmt.Printf("a=%v\n", a)

	a.Print0()
	a.Set(100000)

	a.Print0()
/*
	var b int64
	b.Print0()
	*/
}