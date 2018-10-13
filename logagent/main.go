package main

import (
	"strings"
	"github.com/gostudy03/oconfig"
	"github.com/gostudy03/logagent/kafka"
	"github.com/gostudy03/logagent/tailf"
	"fmt"
)

var (
	appConfig AppConfig
)

type AppConfig struct {
	KafkaConf KafkaConfig `ini:"kafka"`
	LogsConf  LogsConfig `ini:"log_conf"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	QueueSize int  `ini:"queue_size"`
}

type LogsConfig struct {
	LogFilenames string `ini:"log_filenames"`
}

func initConfig(filename string) (err error) {

	err = oconfig.UnMarshalFile(filename, &appConfig)
	if err != nil {
		return
	}

	fmt.Printf("read config succ, config:%#v\n", appConfig)
	return
}

func run() (err error){

	//不断从tailf里面读取日志数据，然后通过kakfa发送
	for {
		//1.从tailf读取数据
		line, err := tailf.ReadLine()
		if err != nil {
			continue
		}

		fmt.Printf("line:%s\n", line.Text)
		msg := &kafka.Message{
			Line: line.Text,
			Topic: "nginx_log",
		}

		err = kafka.SendLog(msg)
		if err != nil {
			fmt.Printf("send log failed, err:%v\n", err)
		}
		fmt.Printf("send to kafka succ\n")
	}

	return
}

func main() {
	err := initConfig("./conf/config.ini")
	if err != nil {
		panic(fmt.Sprintf("init config failed, err:%v", err))
	}

	address := strings.Split(appConfig.KafkaConf.Address, ",")
	err = kafka.Init(address, appConfig.KafkaConf.QueueSize)
	if err != nil {
		panic(fmt.Sprintf("init kafka client failed, err:%v", err))
	}
	fmt.Printf("init kafka succ\n")

	err = tailf.Init(appConfig.LogsConf.LogFilenames)
	if err != nil {
		panic(fmt.Sprintf("init tailf client failed, err:%v", err))
	}

	fmt.Printf("init tailf succ\n")
	err = run()
	if err != nil {
		fmt.Printf("run failed, err:%v\n", err)
		return
	}
	fmt.Printf("run finished\n")
}