package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gostudy03/xlog"
)

var (
	client sarama.SyncProducer
	msgChan chan *Message
)

type Message struct {
	Data string
	Topic string
}

func Init(address []string, chanSize int) (err error){

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		xlog.LogError("producer close, err:", err)
		return
	}

	msgChan = make(chan *Message, chanSize)
	go sendKafka()
	return
	/*
	defer client.Close()

	for {
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(time.Second)
	}
	*/
}

func sendKafka() {
	for msg := range msgChan {

		kafkaMsg := &sarama.ProducerMessage{}
		kafkaMsg.Topic =  msg.Topic
		kafkaMsg.Value = sarama.StringEncoder(msg.Data)
	

		pid, offset, err := client.SendMessage(kafkaMsg)
		if err != nil {
			xlog.LogError("send message failed,", err)
			continue
		}
		xlog.LogDebug("pid:%v offset:%v", pid, offset)
	}
}

func SendLog(msg *Message) (err error) {

	if len(msg.Data) == 0 {
		return
	}

	select {
	case msgChan <- msg:
	default:
		err = fmt.Errorf("chan is full")
	}

	return
}