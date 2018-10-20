package main

import (
	"sync"

	"fmt"
	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {

	
	consumer, err := sarama.NewConsumer([]string{"localhost:2181"}, nil)
	if err != nil {
		fmt.Println("consumer close, err:", err)
		return
	}

	fmt.Printf("connect succ\n")
	partitions, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Printf("get partition failed, err:%v\n", err)
		return
	}

	for _, p := range partitions {
		pc, err := consumer.ConsumePartition("nginx_log", p, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("err:%v\n", err)
			continue
		}
		wg.Add(1)
		go func()  {
			messageChan := pc.Messages()
			for m := range messageChan {
				fmt.Printf("message:%v, text:%v\n", m, string(m.Value))
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
