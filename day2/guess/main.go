package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number int  
	/*
	for i := 0; i < 10; i++ {
		number = rand.Intn(100)
		fmt.Printf("number:%d\n", number)
	}
	*/
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(100)
	fmt.Printf("请猜一个数字，数字的范围:[0-100)\n")
	for {
		var input int
		fmt.Scanf("%d\n", &input)
		var flag bool = false
		switch {
		case number > input:
			fmt.Printf("你输入的数字太小\n")
		case number == input:
			fmt.Printf("恭喜你，猜对了\n")
			flag = true
		case number < input:
			fmt.Printf("你输入的数字太大\n")
		}

		if flag {
			break
		}
	}
}
