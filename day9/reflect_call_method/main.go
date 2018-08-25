package main

import (
	"fmt"
	"reflect"
)


type User struct {
	Name string  `json:"name"`
	Age int
	Sex string
}

func (u *User) Print() {
	fmt.Printf("name:%s age:%d sex:%s\n", u.Name, u.Age, u.Sex)
}

func (u *User) SetName(name string) {
	u.Name = name
}

//1. 获取a的类型
//2. 我要动态改变a里面存的值
//3. 如果a里面存储的是一个结构体，那可以通过反射获取结构体中的字段信息以及调用结构体里面的方法
func TestValue(a interface{}) {


	//调用没有参数的方法
	v := reflect.ValueOf(a)
	methodNum := v.NumMethod()
	fmt.Printf("method:%v\n", methodNum)
	
	m := v.MethodByName("Print")
	
	var args []reflect.Value
	m.Call(args)

	//调用有参数的方法
	v = reflect.ValueOf(a)
	m = v.MethodByName("SetName")
	
	args = args[0:0]
	args = append(args, reflect.ValueOf("hello world"))
	m.Call(args)
}

func main()  {
	
	var user User
	user.Name = "xxx"
	user.Age = 100
	user.Sex = "man"
	TestValue(&user)
	fmt.Printf("user:%#v\n", user)

	/*
	func1 := func(a int) {
		fmt.Print(a)
	}
	fmt.Printf("%p\n", func1)
	//func1(100)
	*/
}