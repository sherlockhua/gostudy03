package main


import (
	"fmt"
)

func main() {
	var arr [10]int
	for i := 0; i < len(arr);i++ {
		arr[i] = i
	}

	var sum int = 12
	for i := 0; i < len(arr);i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i] + arr[j] == sum) {
				fmt.Printf("i=%d j=%d\n", i, j)
			}
		}
	}
	
}