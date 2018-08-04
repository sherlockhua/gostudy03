package main


import (
	"os"
	"fmt"
	"io"
	"compress/gzip"
	_"bufio"
)

func main() {
	fmt.Printf("start run...\n")

	filename := "c:/tmp/xttr.log.gz"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}
	fmt.Printf("start0 read file\n")
	defer file.Close()
	/*
	defer func() {
		file.Close()
	}()
		*/
		
	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("gzip read failed, err:%v\n", err)
		return
	}
	
	//reader := bufio.NewReader(file)
	var content []byte
	var buf[400]byte
	for {
		//reader.Read
		fmt.Printf("start read file\n")
		n, err := reader.Read(buf[:])
		fmt.Printf("read %d  err:%v\n", n,  err)
		if err != nil && err != io.EOF {
			fmt.Printf("read %s failed, err:%v\n", filename, err)
			return
		}

		if n > 0 {
			validBuf := buf[0:n]
			//fmt.Printf("%s\n", string(validBuf))
			content = append(content, validBuf...)
		}
		//读到文件末尾了，文件已经读取完毕，Read方法会返回一个io.EOF错误。
		if err == io.EOF {
			break
		}
	}

	//fmt.Printf("content:%s\n", content)
}