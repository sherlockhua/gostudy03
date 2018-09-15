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

func TestValueConvert(str string){

	//1. string -> interface{} -> reflect.Value
	//var s1 interface{}
	//s1 = str
	value := reflect.ValueOf(str)

	//2. reflect.Value ->interface{} ->string
	s2 := value.Interface()
	str2 := s2.(string)
	fmt.Printf("s:%s\n", str2)
}

func typeConvert() {
	
	var c string = "hello"
	var a int = 100
	var b interface{}

	b = a
	a1, ok := b.(string)
	if !ok {
		fmt.Printf("b is not string\n")
	}
	fmt.Printf("a1=%s\n", a1)	

	b = c
	c1 := b.(string)
	fmt.Printf("c1=%s\n", c1)	

}

func main()  {
	
	/*
	var user User
	user.Name = "xxx"
	user.Age = 100
	user.Sex = "man"
	TestValue(user)
	fmt.Printf("user:%#v\n", user)
	*/
	//var a string = "hello"
	//TestValueConvert(a)
	typeConvert()
}