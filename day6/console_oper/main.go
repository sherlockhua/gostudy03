package main


import (
	"os"
	"fmt"
	
	"bufio"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hello world")
	writer.Flush()

	fmt.Printf("hello world")

	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("read from console failed, err:%v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "data:%s\n", data)
}