package main


import (
	"flag"
	"fmt"
)


func main() {
	var num int
	var mode string

	flag.IntVar(&num, "num", 16, "-num the password length")
	flag.StringVar(&mode, "mode", "mix", "-mode the password generate mode")

	flag.Parse()

	fmt.Printf("num:%d mode:%s\n", num, mode)
}