package main



import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
)

func greet(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	fmt.Fprintf(w, "greet Hello World! word:%s", word)
}


func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}



func userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.method:%s\n", r.Method)
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./login.html")
		if err != nil {
			http.Redirect(w, r, "/404.html", http.StatusNotFound)
			return
		}

		w.Write(data)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "admin" {
			fmt.Fprintf(w, "login success")
		} else {
			fmt.Fprintf(w, "login failed")
		} 
	}
}


func main() {
	http.HandleFunc("/index", greet)
	http.HandleFunc("/user/info", userInfo)
	http.HandleFunc("/user/login", userLogin)
	//http.HandleFunc("/user/info", userInfo)
	http.ListenAndServe(":8080", nil)
}