package main


import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name string
	Age int
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

func (a *Animal) SetAge(age int) {
	a.Age = age
}

func (a *Animal) Print() {
	fmt.Printf("a.name=%s a.age=%d\n", a.Name, a.Age)
}

type Birds struct {
	*Animal
}

type Dog struct {
	*Animal
}

func (b *Birds) Fly() {
	fmt.Printf("name %s is fly\n", b.Name)
}

func main () {
	var b *Birds = &Birds{
		Animal: &Animal{},
	}
	b.SetName("bird")
	b.SetAge(18)

	b.Fly()

	data, err := json.Marshal(b)
	fmt.Printf("marshal result:%s err:%v\n", string(data), err)

	var c Birds
	err = json.Unmarshal(data, &c)
	fmt.Printf("c:%#v, err:%v\n", c.Animal, err)
}