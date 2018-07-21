package main


import (
	"fmt"
)

func isPrime(n int) bool {
	var flag = true
	for j := 2; j < n; j++ {
		if (n % j == 0) {
			flag = false
			break
		}
	}
	
	return flag
}

func main() {
	var n int
	fmt.Printf("please input n:\n")
	fmt.Scanf("%d", &n)
	for i := 2; i < n; i++ {
		if isPrime(i) {
			fmt.Printf("%d\n", i)
		}
	}
}