package main



import (
	"fmt"
)

type User struct {
	Name string
	Sex string
	Age int
	AvatarUrl string
}

func NewUser(name, sex string, age int, url string) User {
	var user User
	user.Name = name
	user.AvatarUrl = url
	user.Age = age
	user.Sex = sex
	return user
}

func main() {
	var user User
	user.Age = 100
	user.Sex = "male"
	user.Name = "jim"
	user.AvatarUrl = "https://baidu.com/xx.jpg"

	fmt.Printf("user.name:%s user.age:%d\n", user.Name, user.Age)

	user02 := User{
		Name: "user02",
		Age: 18,
		Sex: "male",
	}

	fmt.Printf("user02.name:%s user02.age:%d\n", user02.Name, user02.Age)
	user03 := User{}
	fmt.Printf("user03:%#v\n", user03)

	fmt.Printf("name addr:%p\n", &user03.Name)
	fmt.Printf("Sex addr:%p\n", &user03.Sex)
	fmt.Printf("Age addr:%p\n", &user03.Age)
	fmt.Printf("AvatarUrlname addr:%p\n", &user03.AvatarUrl)

	user04 := NewUser("user04", "female", 18, "xxxxx")
	fmt.Printf("user04:%#v\n", user04)
}
