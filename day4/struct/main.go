package main


import (
	"fmt"
)

//1. 定义个Student的类型
type Student struct {
	Name string
	Age int
}

//定义了一个值类型为Student的GetName方法
func (s Student) GetName() string {
	return s.Name
}

//定义了一个指针类型为Student的SetName方法
func (s *Student) SetName(name string) {
	s.Name = name
}

func main() {
	//2.定义一个类型为Student的变量
	var s1 Student = Student{
		Name: "s1",
		Age: 18,
	}

	name := s1.GetName()
	fmt.Printf("name=%s\n", name)

	//(&s1).SetName("s2")
	s1.SetName("s2")
	name = s1.GetName()
	fmt.Printf("name=%s\n", name)

	s1.Print0()
}