package etcd

import (
	"fmt"
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
	watchKey string
	dataCh chan []*common.CollectConfig
}

var (
	etcdClient *EtcdClient
)

func Init(address []string, watchKey string)(err error) {
	etcdClient = &EtcdClient {
		address: address,
		watchKey:watchKey,
		dataCh: make(chan []*common.CollectConfig),
	}

	etcdClient.client, err = clientv3.New(clientv3.Config{
		Endpoints: address,
		DialTimeout:3 * time.Second,
	})

	if err != nil {
		xlog.LogError("create etcd client failed, err:%v, address:%v", err, address)
		return
	}

	go etcdClient.watch()
	return
}

func (e *EtcdClient) watch() {
	for {
		resultCh := e.client.Watch(context.Background(), e.watchKey)
		xlog.LogDebug("wacth return, resultCh:%v", resultCh)
		for v := range resultCh {
			xlog.LogDebug("wacth return, v:%v", v)
			if v.Err() != nil {
				xlog.LogError("watch:%s failed, err:%v\n",e.watchKey, v.Err())
				continue;
			}

			for _, event := range v.Events {
				xlog.LogDebug("event_type:%v key:%v val:%v\n", event.Type, event.Kv.Key, string(event.Kv.Value))
				var conf []*common.CollectConfig
				if event.Type == clientv3.EventTypeDelete {
					e.dataCh <- conf
					continue
				}

				err := json.Unmarshal(event.Kv.Value, &conf)
				if err != nil {
					xlog.LogWarn("unmarshal failed, key:%s val:%s", e.watchKey, string(event.Kv.Value))
					continue
				}

				e.dataCh <- conf
			}
		}
	}
}

func Watch() <- chan []*common.CollectConfig{

	return etcdClient.dataCh
}

func GetCollectSystemInfoConfig(key string) (conf *common.CollectSystemInfoConfig, err error) {
	resp , err := etcdClient.client.Get(context.Background(), key)
	if err != nil {
		xlog.LogError("get key:%s from etcd failed, err:%v", key, err)
		return
	}

	if (len(resp.Kvs) == 0) {
		xlog.LogError("get key:%s from etcd failed, len(resp.kvs)=0", key)
		err = fmt.Errorf("not found value of %s", key)
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

func GetConfig(key string) (conf []*common.CollectConfig, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer func() {
		cancel()
	}()

	resp , err := etcdClient.client.Get(ctx, key)
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
