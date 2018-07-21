package main


import (
	"fmt"
	"time"
)

type User struct {
	name string
	s1 [10000000]int64
	s2 [10000000]int64
	s3 [10000000]int64
	s4 [10000000]int64
}

func (u *User) Set() {
	for i := 0; i < len(u.s1); i++{
		u.s1[i] = 1
		u.s2[i] = 1
		u.s3[i] = 1
		u.s4[i] = 1
	}
}

func main() {
	//var u *User = new(User) 
	var u *User = &User{
		name:"user01", 
	}
/*
	var u2 User = User {
		name :"user02",
	}
*/
	start := time.Now().UnixNano()
	u.Set()
	end := time.Now().UnixNano()

	fmt.Printf("cost:%d ns", (end - start)/1000)
}