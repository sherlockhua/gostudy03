package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return x, nil
}

func main() {
	fmt.Println(Sqrt(2))
	_, err := Sqrt(-2)
	if err != nil {
	switch err.(type) {
	case ErrNegativeSqrt:
		fmt.Printf("ErrNegativeSqrt\n")
	default:

	}
}
}
