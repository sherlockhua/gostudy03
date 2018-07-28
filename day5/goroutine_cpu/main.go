package main


import (
	"time"
	"runtime"
	"fmt"
)

func main() {
	cpu := runtime.NumCPU()
	//runtime.GOMAXPROCS(1)

	for i := 0; i < 8; i++ {
		go func() {
			for {

			}
		}()
	}
	fmt.Printf("%d\n", cpu)
	time.Sleep(15*time.Second)
}