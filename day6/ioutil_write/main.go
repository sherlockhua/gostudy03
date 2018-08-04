package main


import (
	"io/ioutil"
//	"os"
//	"fmt"
	
)

func main() {
	filename := "D:\\project\\src\\github.com\\gostudy03\\day5\\atomic\\main.go"
	str := "dkfslfjdsklfjlskjflsjflsjflsjflks"
	ioutil.WriteFile(filename, []byte(str), 0755)
}