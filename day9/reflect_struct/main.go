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

//1. 获取a的类型
//2. 我要动态改变a里面存的值
//3. 如果a里面存储的是一个结构体，那可以通过反射获取结构体中的字段信息以及调用结构体里面的方法
func TestValue(a interface{}) {

	v := reflect.ValueOf(a)
	t := v.Type()
	switch t.Kind()  {
	case reflect.Struct:
		fieldNum := t.NumField()
		fmt.Printf("field num:%d\n", fieldNum)
		for i := 0; i <fieldNum; i++{
			field := t.Field(i)
			vField := v.Field(i)
			
			fmt.Printf("field[%d] name:%s, json key:%s, val:%v\n",
				 i, field.Name, field.Tag.Get("json"), vField.Interface())
		}

	}
}

func main()  {
	
	var user User
	user.Name = "xxx"
	user.Age = 100
	user.Sex = "man"
	TestValue(user)
	fmt.Printf("user:%#v\n", user)
}