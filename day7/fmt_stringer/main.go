package main


import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name string
	Age int
}

func (s *Student)String()string {
	data, _ := json.Marshal(s)
	return string(data)
}

func main() {
	var a = &Student{
		Name: "hell",
		Age:12,
	}
	fmt.Printf("a = %v\n", a)
}