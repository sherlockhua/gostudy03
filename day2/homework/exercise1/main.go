package main


import (
	"math/rand"
	"fmt"
)

func main() {
	var arr [10]int
	for i := 0; i < len(arr);i++ {
		arr[i] = rand.Intn(10000)
	}

	var sum int
	for i := 0; i < len(arr);i++ {
		sum = sum + arr[i]
	}
	fmt.Println("sum:", sum)
}