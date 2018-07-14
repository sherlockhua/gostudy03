package main

import (
	"sort"
	"fmt"
)

func initMap1() {
	var user map[string]int = make(map[string]int, 5000)
	user["abc"] = 38
	user["user01"] = 10000
	user["user02"] = 10002

	fmt.Printf("user:%v\n", user["userxxxxx"])
}

func initMap2() {
	var m map[string]int
	m = map[string]int{
		"user01": 10000,
		"user02": 20000,
	}

	m["user03"] = 3838
	fmt.Printf("user:%#v\n", m)
}

var whiteUser map[int]bool = map[int]bool{
	32423:   true,
	3483943: true,
	1:       true,
}

func isWhiteUser(userId int) bool {
	_, ok := whiteUser[userId]
	return ok
}

func testWhiteUser() {
	userId := 3483943
	if isWhiteUser(userId) {
		fmt.Printf("is white user:%v\n", userId)
	} else {
		fmt.Printf("is normal user:%v\n", userId)
	}
}

func transverse() {
	var m map[string]int
	m = map[string]int{
		"user01": 10000,
		"user02": 20000,
	}

	m["user03"] = 3838

	for key, value := range m {
		fmt.Printf("key:%s value:%d\n", key, value)
	}
}

func testDelete() {
	var m map[string]int
	m = map[string]int{
		"user01": 10000,
		"user02": 20000,
	}

	m["user03"] = 3838

	for key, value := range m {
		fmt.Printf("key:%s value:%d\n", key, value)
	}

	fmt.Println("........")
	delete(m, "user01")

	for key, value := range m {
		fmt.Printf("key:%s value:%d\n", key, value)
	}
}

func testMapCopy() {
	a := map[string]int{
		"steve ": 12000,
		"jamie ": 15000,
	}
	a["mike"] = 9000
	fmt.Println("origin map", a)
	b := a
	b["mike"] = 18000
	fmt.Println(" a map changed", a)
}

func testMapSort() {
	a := map[string]int{
		"steve ": 12000,
		"jamie ": 15000,
	}
	a["mike"] = 9000

	var keys []string
	for key, _ := range a {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("key:%s val:%v\n", key, a[key])
	}
}

func testMapSlice() {
	var s []map[string]int
	s = make([]map[string]int, 5)

	for k, v := range s {
		fmt.Printf("index:%d val:%v\n", k, v)
	}

	s[0] = make(map[string]int, 16)
	s[0]["abc"] = 100

	for key, val := range s[0] {
		fmt.Printf("key:%s val:%v\n", key, val)
	}
}

func main() {
	//initMap1()
	//initMap2()
	//testWhiteUser()
	//transverse()
	//testDelete()
	//testMapCopy()
	//testMapSort()
	testMapSlice()
}
