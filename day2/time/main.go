package main

import (
	"fmt"
	"time"
)

func testTimer(){
	//NewTimer定时器，只触发一次
	timer := time.NewTimer(time.Second)
	
	for v := range timer.C {
		fmt.Printf("time:%v\n", v)
		//必须reset，否则会一直堵塞
		timer.Reset(time.Second)
	}
}


func testTicker(){
	//NewTicker定时器，每隔触发一次
	timer := time.NewTicker(time.Second)
	
	for v := range timer.C {
		fmt.Printf("time:%v\n", v)
	}
}

func timestampToTime(timestamp int64) {
	t := time.Unix(timestamp, 0)

	fmt.Printf("convert timestamp to time:%v\n", t)
}


func main() {
	//var now time.Time
	//now = time.Now()
	now := time.Now()
	fmt.Printf("current time is %v\n", now)
	fmt.Printf("%d-%d-%d %02d:%02d:%02d\n", 
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("timestamp:%d ns:%d\n", now.Unix(), now.UnixNano())

	timestampToTime(now.Unix())
	 //testTimer()
	 //testTicker()
	time.Sleep(time.Minute)
}