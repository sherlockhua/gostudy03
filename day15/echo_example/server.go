package main

import (
	"time"
	"flag"
	"log"
	"net/http"

	"html/template"

	"github.com/gorilla/websocket"
)

var (
	homeTemplate *template.Template
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		/*
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
		*/
		c.WriteMessage(websocket.TextMessage, []byte("this is server push"))
		time.Sleep(time.Second)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func initTemplate() (err error) {
	//homeTemplate.ParseFiles("./view/index.html")
	homeTemplate, err = template.ParseFiles("./view/index.html")
	if err != nil {
		log.Fatalf("parse index.html failed, err:%v", err)
	}
	return
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	initTemplate()
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
