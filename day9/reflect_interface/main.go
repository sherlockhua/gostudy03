package main

import (
	"fmt"
	"reflect"
)

//1. 获取a的类型
//2. 我要动态改变a里面存的值
//3. 如果a里面存储的是一个结构体，那可以通过反射获取结构体中的字段信息以及调用结构体里面的方法

func TestType(a interface{}) {

	t := reflect.TypeOf(a)
	fmt.Printf("t = %v\n", t)

	kind := t.Kind()
	switch kind {
	case reflect.Int:
		fmt.Printf("a is int\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	}
}


func TestValue(a interface{}) {

	v := reflect.ValueOf(a)
	//v.Type() 和 reflect.TypeOf(a)的功能是一样

	t := v.Type()
	switch t.Kind()  {
	case reflect.Int:
		v.SetInt(1000)
	case reflect.String:
		v.SetString("xxxxx")
	case reflect.Ptr:
		//t1 := v.Elem().Type()
		e := v.Elem()
		t1 := e.Type()
		switch (t1.Kind()) {
			case reflect.Int:
				v.Elem().SetInt(1000)
				fmt.Printf("ptr is int\n")
			case reflect.String:
				v.Elem().SetString("hello")
				fmt.Printf("ptr is string\n")
		}
		fmt.Printf("a is point type\n")
	}
}

func main()  {
	var a int
	TestType(a)
	var b string
	TestType(b)
	fmt.Printf("a=%v\n", a)

	TestValue(&a)
	TestValue(&b)

	fmt.Printf("a=%v b = %v \n", a, b)
}