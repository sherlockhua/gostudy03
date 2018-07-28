package main

import (
	"time"
)


type Book struct {
	ID string
	Name string
	Num uint32
	Author string
	Publish time.Time
}