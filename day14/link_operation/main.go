package main


import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p *Person)SetName(name string) *Person{
	p.Name = name
	return p
}

func (p *Person) SetAge(age int) *Person{
	p.Age = age
	return p 
}

func (p *Person) Print() {
	fmt.Printf("name:%#v age:%#v\n", p.Name, p.Age)
}

func main() {
	p := &Person{}
	p.SetName("zhangsan").SetAge(28).Print()

}