package main


import (
	"os"
	"fmt"
	
)

func isFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func main() {
	filename := "D:\\project\\src\\github.com\\gostudy03\\day5\\atomic\\main.go"

	var file *os.File
	var err error
	if isFileExists(filename) {
		//mac机器
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0755)
	} else {
		file, err = os.Create(filename)
	}

	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}

	defer file.Close()
	/*
	n , err := file.WriteString("hello world")
	if err != nil {
		fmt.Printf("write failed, err:%v\n", err)
		return
	}
	*/
	fmt.Fprintf(file, "%d %d is good", 100, 300)
	//fmt.Printf("write %d succ\n", n)
}