package main
import (
	"fmt"
	"time"
)

func main() {
	//var now time.Time
	//now = time.Now()
	now := time.Now()
	fmt.Printf("%d-%d-%d %02d:%02d:%02d\n", 
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		//2006/02/01 15:04:05 是go的诞生时间，必须写正确
	str := now.Format("2006-02-01 15:04:05")
	fmt.Printf("str:%s\n", str)
}