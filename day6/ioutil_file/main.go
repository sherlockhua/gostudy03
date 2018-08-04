package main


import (
	"io/ioutil"
	"fmt"
)

func main() {
	filename := "D:\\project\\src\\github.com\\gostudy03\\day5\\atomic\\main.go"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file %s failed, err:%v\n", filename)
		return
	}
	fmt.Printf("content:%s\n", string(content))
}