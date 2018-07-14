package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age int
}

type People struct {
	Animal
	Sex string
}

func main () {
	var p People
	p.Age = 18
	p.Name = "user01"
	p.Sex = "male"

	fmt.Printf("people:%#v\n", p)
}