package main

import (
	"fmt"
)

func main(){
	var a [3][2]int = [3][2]int {
		{1,2},
		{2,3},
		{3,4},
	}
	for index, row := range a {
		fmt.Printf("row:%d value:%v\n", index, row)
	}
	for index, row := range a {
		for col, value := range row {
			fmt.Printf("a[%d][%d]=%d\n", index, col, value)
		}
	}
}