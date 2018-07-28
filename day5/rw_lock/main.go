package main


import (
	"time"
	"fmt"
	"sync"
)

var rwlock sync.RWMutex
var mlock sync.Mutex
var wg sync.WaitGroup
var count int

func writer() {
	for i := 0; i < 1000; i++{
		// 加写锁
		rwlock.Lock()
		count++
		time.Sleep(10*time.Millisecond)
		// 释放写锁
		rwlock.Unlock()
	}
	wg.Done()
}

func reader() {
	for i := 0; i < 1000; i++{
		// 加读锁
		rwlock.RLock()
		_ = count
		//fmt.Printf("count=%d\n", count)
		time.Sleep(1*time.Millisecond)
		// 释放读锁
		rwlock.RUnlock()
	}
	wg.Done()
}


func writer_mutex() {
	for i := 0; i < 1000; i++{
		mlock.Lock()
		count++
		time.Sleep(10*time.Millisecond)
		mlock.Unlock()
	}
	wg.Done()
}

func reader_mutex() {
	for i := 0; i < 1000; i++{
		mlock.Lock()
		//fmt.Printf("count=%d\n", count)
		time.Sleep(1*time.Millisecond)
		mlock.Unlock()
	}
	wg.Done()
}


func main(){

	start := time.Now().UnixNano()
	wg.Add(1)
	//go writer()
	go writer_mutex()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//go reader()
		go reader_mutex()
	}

	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start)/1000/1000/1000
	fmt.Printf("cost %d s\n", cost)
}