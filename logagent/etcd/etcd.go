package etcd

import (
	"context"
	"time"
	"go.etcd.io/etcd/clientv3"
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/logagent/common"
	"encoding/json"
)

type EtcdClient struct {
	client *clientv3.Client
	address []string
}

var (
	etcdClient *EtcdClient
)

func Init(address []string)(err error) {
	etcdClient = &EtcdClient {
		address: address,
	}

	etcdClient.client, err = clientv3.New(clientv3.Config{
		Endpoints: address,
		DialTimeout:3 * time.Second,
	})

	if err != nil {
		xlog.LogError("create etcd client failed, err:%v, address:%v", err, address)
		return
	}

	return
}

func GetConfig(key string) (conf []*common.CollectConfig, err error) {

	resp , err := etcdClient.client.Get(context.Background(), key)
	if err != nil {
		xlog.LogError("get key:%s from etcd failed, err:%v", key, err)
		return
	}

	if (len(resp.Kvs) == 0) {
		xlog.LogError("get key:%s from etcd failed, len(resp.kvs)=0", key)
		return
	}
	
	keyVals := resp.Kvs[0]
	xlog.LogDebug("get key:%s from etcd succ, key:%v val:%v", key, keyVals.Key, keyVals.Value)

	err = json.Unmarshal(keyVals.Value, &conf)
	if err != nil {
		xlog.LogError("unmarshal failed, data:%v", string(keyVals.Value))
		return
	}

	xlog.LogDebug("get config from etcd succ, conf:%#v", conf)
	return
}
