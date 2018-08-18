package main

import (
	"html/template"
	"fmt"
	"net/http"
)

var (
	t *template.Template
)

type User struct {
	Name string
	Age int
}
func initTemplate() (err error){
	t, err = template.ParseFiles("./index.html")
	if err != nil {
		fmt.Printf("load template failed,err:%v\n", err)
		return
	}
	return
}


func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	var user User = User{
		Name:"user01", 
		Age:10,
	}

	t.Execute(w, user)
}

func main() {
	
	err := initTemplate()
	if err != nil {
		return
	}

	http.HandleFunc("/user/info", handleUserInfo)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
}