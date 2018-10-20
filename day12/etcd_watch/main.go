package main

import (
	"context"
	"time"
	"go.etcd.io/etcd/clientv3"
	"fmt"
)


func main()  {

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.20.200:2379"},
		DialTimeout:3 * time.Second,
	})

	defer client.Close()

	fmt.Printf("conn succ\n")
	for {
		resultCh := client.Watch(context.Background(), "/logagent/", clientv3.WithPrefix())
		fmt.Printf("wacth return, resultCh:%v\n", resultCh)
		for v := range resultCh {
			fmt.Printf("wacth return, v:%v\n", v)
			if v.Err() != nil {
				fmt.Printf("watch failed, err:%v\n", err)
				continue;
			}

			for _, e := range v.Events {
				fmt.Printf("event_type:%v key:%v val:%v\n", e.Type, e.Kv.Key, string(e.Kv.Value))
			}
		}
	}
}