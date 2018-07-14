package main


import (
	"sort"
	"math/rand"
	"fmt"
)

func main() {
	var arr [10]int
	for i := 0; i < len(arr);i++ {
		arr[i] = rand.Intn(10000)
	}
	fmt.Printf("arr:%v\n", arr)
	sort.Ints(arr[:])

	fmt.Printf("arr:%v\n", arr)
}