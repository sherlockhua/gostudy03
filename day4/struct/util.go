package main

import (
	"fmt"
)

func (s *Student) Print0(){
	fmt.Printf("student:%#v", *s)
}