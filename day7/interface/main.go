package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Talk()
	Run()
}

type PuRuAnimal interface {
	ChiNai()
}

type AdvanceAnimal interface {
	Animal
	PuRuAnimal
}

type Dog struct {
	name string
}

func (d Dog) Eat() {
	fmt.Printf("%s is eating\n", d.name)
}


func (d Dog) ChiNai() {
	fmt.Printf("%s is ChiNai\n", d.name)
}

func (d Dog) Talk() {
	fmt.Printf("%s is talking\n", d.name)
}

func (d Dog) Run() {
	fmt.Printf("%s is runing\n", d.name)
}

type Pig struct {
	name string
}

func (d *Pig) Eat() {
	fmt.Printf("%s is eating\n", d.name)
}


func (d *Pig) Talk() {
	fmt.Printf("%s is talking\n", d.name)
}

func (d *Pig) Run() {
	fmt.Printf("%s is runing\n", d.name)
}

func Describe(a Animal) {
	/* 有坑
	dog := a.(*Dog)
	dog.Eat()
	*/

	dog, ok := a.(*Dog)
	if !ok {
		fmt.Printf("convert to dog failed\n")
		return
	}
	fmt.Printf("decrible succ\n")
	dog.Run()
	fmt.Printf("decrible succ------\n")
}


func DescribeSwitch(a Animal) {
	fmt.Printf("DescribeSwitch(a) begin\n")
	switch v := a.(type) {
	case *Dog:
		fmt.Printf("v = %v\n", v)
		var dog *Dog
		//v就是断言之后的具体类型
		dog = v
		dog.Run()
	case *Pig:
		pig := a.(*Pig)
		pig.Run()
	}
	fmt.Printf("DescribeSwitch(a) end\n")
}

func main() {
	var dog = &Dog{
		name:"旺财",
	}

	var a Animal
	fmt.Printf("a:%v dog:%v\n", a, dog)
	
	a = dog
	Describe(a)
	DescribeSwitch(a)
	a.Eat()
	a.Run()
	a.Talk()

	var pig = &Pig {
		name:"佩奇",
	}

	a = pig
	a.Eat()
	a.Run()
	a.Talk()
	Describe(a)
	DescribeSwitch(a)
	//dog = a

	var dogVal = Dog{
		name:"来福",
	}

	a = dogVal
	a.Run()

	var b PuRuAnimal
	b = dog
	b.ChiNai()

	var c AdvanceAnimal
	c = dog
	c.Eat()
}
