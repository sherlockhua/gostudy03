package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
	msgChan chan *Message
)

type Message struct {
	Line string
	Topic string
}

func Init(address []string, chanSize int) (err error){

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	msgChan = make(chan *Message, chanSize)
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

func SendLog(msg *Message) (err error) {

	select {
	case msgChan <- msg:
	default:
		err = fmt.Errorf("chan is full")
	}

	return
}