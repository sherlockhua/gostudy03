package main


import (
	"fmt"
)

func calc(op func(args...int)int, op_args...int) int {
	result := op(op_args...)
	fmt.Printf("result = %d\n", result)
	return result
}

func add(args...int)int{
	var sum int
	for i := 0; i < len(args);i++ {
		sum = sum + args[i]
	}
	return sum
}

func main() {
	calc(func(args...int)int{
		var sum int
		for i := 0; i < len(args);i++ {
			sum = sum + args[i]
		}
		return sum
	}, 1,2,3,4,5)

	calc(func(args...int)int{
		var sum int
		for i := 0; i < len(args);i++ {
			sum = sum - args[i]
		}
		return sum
	}, 1,2,3,4,5)

	calc(add, 1,2,3,4,5)
}