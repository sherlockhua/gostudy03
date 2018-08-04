package main


import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("args count:%d\n", len(os.Args))
	for index, v := range os.Args {
		fmt.Printf("args%d, value:%s\n", index, v)
	}
}