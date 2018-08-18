package main

import (
	"os"
	"fmt"
)

type Test struct {
	data string
}

func (t *Test) Write(p []byte) (n int, err error) {
	t.data = string(p)
	return len(p), nil
}


func main() {
	file, _ := os.Create("c:/tmp/c.txt")
	fmt.Fprintf(os.Stdout, "hello world\n")
	fmt.Fprintf(file, "hello world\n")
	/*
	fmt.FPtrintfConsole()
	fmt.FPtrintfFile()
	fmt.FPtrintfNet()
	*/
	var t *Test = &Test{}
	fmt.Fprintf(t, "this is a test inteface:%s", "?akdfkdfjdkfk")

	fmt.Printf("t.data:%s\n", t.data)
}