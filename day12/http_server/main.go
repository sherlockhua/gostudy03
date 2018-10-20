package main


import (
	"time"
	"net/http"
)

func watchHandle(w http.ResponseWriter, r *http.Request) {
	
	var count int
	for {
		time.Sleep(time.Second)
		w.Write([]byte("hello"))
	
		count++
		if count > 5 {
			break
		}
	}
}


func main()  {
	http.HandleFunc("/watch", watchHandle)
	http.ListenAndServe(":8080", nil)
}