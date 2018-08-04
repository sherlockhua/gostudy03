package main


import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func main() {
	filename := "D:\\project\\src\\github.com\\gostudy03\\day5\\atomic\\main.go"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}

	defer file.Close()
	/*
	defer func() {
		file.Close()
	}()
		*/
	reader := bufio.NewReader(file)
	var content []byte
	var buf[4096]byte
	for {
		//reader.Read
		n, err := reader.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Printf("read %s failed, err:%v\n", filename, err)
			return
		}
		//读到文件末尾了，文件已经读取完毕，Read方法会返回一个io.EOF错误。
		if err == io.EOF {
			break
		}

		validBuf := buf[0:n]
		//fmt.Printf("%s\n", string(validBuf))
		content = append(content, validBuf...)
	}

	fmt.Printf("content:%s\n", content)
}