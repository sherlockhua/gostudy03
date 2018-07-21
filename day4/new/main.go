package main

import (
	"fmt"
)

type User struct {
	Name string
	age int
}

func test1() {
	var p *int = new(int)
	*p = 1000
	fmt.Printf("p:%v address:%v\n", *p, p)

	var pUser *User = new(User)
	(*pUser).age = 100
	pUser.Name = "user01"

	fmt.Printf("user:%v\n", *pUser)
}

func test2() {
	var p *[]int = new([]int)
	*p = make([]int, 10)

	(*p)[0] = 100
	(*p)[2] = 100

	fmt.Printf("p:%#v\n", *p)

	var p1 *map[string]int = new(map[string]int)
	*p1 = make(map[string]int, 10)
	(*p1)["key"] = 100
	(*p1)["key2"] = 200

	fmt.Printf("p:%#v\n", *p1)
}

func  main() {
	test2()
}