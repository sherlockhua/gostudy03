package main


import (
	"bufio"
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
	
	file, err = os.Create(filename)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("hello worldldfdsfsfsf")

	writer.Flush()
}