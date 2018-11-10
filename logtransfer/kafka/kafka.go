package kafka

import (
	"github.com/gostudy03/xlog"
	"github.com/Shopify/sarama"
	"encoding/json"
	"github.com/gostudy03/logtransfer/es"
)

var (
	consumer sarama.Consumer
)

func Init(addr string, topic string)(err error) {

	consumer, err = sarama.NewConsumer([]string{addr}, nil)
	if err != nil {
		xlog.LogError("new kafka consumer failed, err:%v", err)
		return
	}

	xlog.LogDebug("connect to kafka succ")

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		xlog.LogError("get partition failed, err:%v", err)
		return
	}

	xlog.LogDebug("get partition succ, partition:%#v", partitions)

	for _, p := range partitions {
		pc, err := consumer.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			xlog.LogError("consumer partition failed, err:%v", err)
			continue
		}
	
		go func()  {
			messageChan := pc.Messages()
			for m := range messageChan {
				xlog.LogDebug("recv from kafka, text:%v\n", m, string(m.Value))
				var msg map[string]interface{} = make(map[string]interface{}, 16)

				err = json.Unmarshal(m.Value, &msg)
				if err != nil {
					xlog.LogError("unmarshal failed, err:%v", err)
					continue
				}
				es.AppendMsg(msg)
				xlog.LogDebug("append msg to es succ, msg:%#v", msg)
			}
		}()
	}

	return
}