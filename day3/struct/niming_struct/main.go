package main


import (
	"fmt"
)

type Address struct {
	City string
	Province string
}

type User struct {
	Name string
	Sex string
	Age int
	AvatarUrl string
	int
	string
	address Address
}


func main() {
	var user User
	user.int = 100
	user.string = "hello"

	user.address.City = "beijing"
	user.address.Province = "beijing"

	fmt.Printf("user:%#v\n", user)

	user01 := User {
		int: 100,
		string: "hello",
		address: Address{
			City:"beijing",
			Province: "beijing",
		},
	}
	fmt.Printf("user01:%#v\n", user01)
}