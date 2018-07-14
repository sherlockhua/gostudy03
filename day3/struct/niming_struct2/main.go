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
	City string
	int
	string
	Address
}


func main() {
	var user User
	user.int = 100
	user.string = "hello"

	user.Address.City = "beijing"
	user.Address.Province = "beijing"

	fmt.Printf("user:%#v\n", user)

	user01 := User {
		int: 100,
		string: "hello",
		Address: Address{
			City:"beijing",
			Province: "beijing",
		},
	}
	fmt.Printf("user01:%#v\n", user01)

	fmt.Printf("city:%s province:%s\n", user01.Address.City, user01.Province)
}