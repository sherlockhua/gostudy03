package main

import (
	"context"
	"time"
	"go.etcd.io/etcd/clientv3"
	"github.com/gostudy03/logagent/common"
	"fmt"
	"encoding/json"
)


func main()  {

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.20.200:2379"},
		DialTimeout:3 * time.Second,
	})

	defer client.Close()

	var logCollectArray []common.CollectConfig
	logCollect := common.CollectConfig{
		Topic:"nginx_log",
		Path:"c:/tmp/a.log",
		ModuleName:"nginx",
	}
	logCollectArray = append(logCollectArray, logCollect)

	logCollect = common.CollectConfig{
		Topic:"nginx_log",
		Path:"c:/tmp/b.log",
		ModuleName:"nginx",
	}

	logCollectArray = append(logCollectArray, logCollect)
	data, err := json.Marshal(logCollectArray)
	if err != nil {
		fmt.Printf("marshal failed, conf:%#v\n", logCollectArray)
		return
	}
	_, err = client.Put(context.Background(), "/logagent/conf", string(data))
	if err != nil {
		fmt.Printf("put failed, err:%v\n", err)
	}
}