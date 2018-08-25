package main

import (
	"fmt"
	"github.com/gostudy03/day9/reflect_json"
)


/*
{
	"name": "xxx",
	"Age":100,
	"Sex": "xx"
}
*/
type User struct {
	Name string  `json:"name"`
	Age int
	Sex string
}

func main() {
	var a string = "hello world"
	jsonStr := reflect_json.Marshal(a)
	fmt.Printf(jsonStr)

	var user User 
	user.Age = 10900
	user.Name = "user01"
	user.Sex = "man"

	jsonStr = reflect_json.Marshal(user)
	fmt.Printf("user marshal:%s\n", jsonStr)
}