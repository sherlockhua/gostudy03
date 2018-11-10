package es

import (
	"context"
	"github.com/gostudy03/xlog"
	elastic "github.com/olivere/elastic"
)

type ESCient struct {
	client *elastic.Client
	index string
	threadNum int
	queueSize int
	queue chan interface{}
}

var (
	esClient *ESCient = &ESCient{}
)

func Init(addr string, index string, threadNum, queueSize int)(err error) {

	xlog.LogDebug("init es addr:%s index:%s thread_num:%d queue_size:%d",
		addr, index, threadNum, queueSize)
	client, err := elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		xlog.LogError("init es client failed, err:%v", err)
		return
	}

	esClient.client = client
	esClient.index  = index
	esClient.threadNum = threadNum
	esClient.queueSize = queueSize
	esClient.queue = make(chan interface{}, queueSize)

	for i := 0; i < threadNum; i++ {
		go insertES()
	}
	return
}

func AppendMsg(msg interface{}) {

	xlog.LogDebug("append msg to es queue, msg:%#v", msg)
	esClient.queue <- msg
}

func insertES() {
	for data := range esClient.queue {
		_, err := esClient.client.Index().
			Index(esClient.index).
			Type(esClient.index).BodyJson(data).Do(context.Background())
		if err != nil {
			xlog.LogError("do insert es failed, err:%v", err)
			continue
		}
	}
}