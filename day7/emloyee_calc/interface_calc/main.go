package main

import (
	"fmt"
)

type Employee interface {
	Calc() float32
} 


type Developer struct {
	Name string
	Base float32
}

func (d *Developer) Calc() float32 {
	return d.Base
}

type PM struct {
	Name string
	Base float32
	Option float32
}


func (p *PM) Calc() float32 {
	return p.Base + p.Option
}

type YY struct {
	Name string
	Base float32
	Option float32
	Rate float32 //0.6 ~ 3
}

func (p *YY) Calc() float32 {
	return p.Base + p.Option * p.Rate
}



type EmployeeMgr struct {
	employeeList []Employee
}

func (e *EmployeeMgr) Calc() float32 {
	var sum float32
	for _, v := range e.employeeList {
		sum += v.Calc()
	}

	return sum
}

func (e *EmployeeMgr) AddEmpoyee(d Employee) {
	e.employeeList = append(e.employeeList, d)
}


func main () {
	var e = &EmployeeMgr{}

	dev := &Developer{
		Name: "develop",
		Base: 10000,
	}
	e.AddEmpoyee(dev)

	pm := &PM{
		Name: "develop",
		Base: 10000,
		Option: 12000,
	}
	e.AddEmpoyee(pm)

	yy := &YY{
		Name: "develop",
		Base: 10000,
		Option: 12000,
		Rate: 1.2,
	}
	e.AddEmpoyee(yy)

	sum := e.Calc()
	fmt.Printf("sum:%f\n", sum)
}