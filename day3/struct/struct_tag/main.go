package main


import (
	"fmt"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Sex string `json:"sex"`
	Age int      `json:"age"`
	AvatarUrl string `json:"avatar_url"`
}


func main(){

	var user User
	user.Age = 100
	user.Sex = "male"
	user.Name = "jim"
	user.AvatarUrl = "https://baidu.com/xx.jpg"

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("json:%v\n", string(data))
}