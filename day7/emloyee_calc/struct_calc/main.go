package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Base int
}

func (d *Developer) Calc() int {
	return d.Base
}

type PM struct {
	Name string
	Base int
	Option int
}


func (p *PM) Calc() int {
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
	devList []*Developer
	pmList []*PM
	yyList []*YY
}

func (e *EmployeeMgr) Calc() float32 {
	var sum float32
	for _, v := range e.devList {
		sum += float32(v.Calc())
	}

	for _, v := range e.pmList {
		sum += float32(v.Calc())
	}

	for _, v := range e.yyList {
		sum += float32(v.Calc())
	}

	return sum
}

func (e *EmployeeMgr) AddDev(d *Developer) {
	e.devList = append(e.devList, d)
}


func (e *EmployeeMgr) AddPM(d *PM) {
	e.pmList = append(e.pmList, d)
}


func (e *EmployeeMgr) AddYY(d *YY) {
	e.yyList = append(e.yyList, d)
}

func main () {
	var e = &EmployeeMgr{}

	dev := &Developer{
		Name: "develop",
		Base: 10000,
	}
	e.AddDev(dev)

	pm := &PM{
		Name: "develop",
		Base: 10000,
		Option: 12000,
	}
	e.AddPM(pm)

	yy := &YY{
		Name: "develop",
		Base: 10000,
		Option: 12000,
		Rate: 1.2,
	}
	e.AddYY(yy)

	sum := e.Calc()
	fmt.Printf("sum:%f\n", sum)
}