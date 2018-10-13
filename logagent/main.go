package main

import (
	"strings"
	"github.com/gostudy03/oconfig"
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/logagent/kafka"
	"github.com/gostudy03/logagent/tailf"
	"fmt"
)

var (
	appConfig AppConfig
)

type AppConfig struct {
	KafkaConf KafkaConfig `ini:"kafka"`
	CollectLogConf  CollectLogConfig `ini:"collect_log_conf"`
	LogConf LogConfig `ini:"logs"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	QueueSize int  `ini:"queue_size"`
}

type CollectLogConfig struct {
	LogFilenames string `ini:"log_filenames"`
}

/*
log_level=debug    
filename=./logs/logagent.log
#console|file
log_type=file
module=logagent
*/
type LogConfig struct {
	LogLevel string `ini:"log_level"`
	Filename string `ini:"filename"`
	LogType string `ini:"log_type"`
	Module string `ini:"module"`
}

func initConfig(filename string) (err error) {

	err = oconfig.UnMarshalFile(filename, &appConfig)
	if err != nil {
		return
	}

	xlog.LogDebug("read config succ, config:%#v", appConfig)
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

		xlog.LogDebug("line:%s", line.Text)
		msg := &kafka.Message{
			Line: line.Text,
			Topic: "nginx_log",
		}

		err = kafka.SendLog(msg)
		if err != nil {
			xlog.LogWarn("send log failed, err:%v\n", err)
		}
		xlog.LogDebug("send to kafka succ\n")
	}

	return
}

func initLog() (err error) {

	var logType int
	var level int

	if appConfig.LogConf.LogType == "console" {
		logType = xlog.XLogTypeConsole
	} else {
		logType = xlog.XLogTypeFile
	}

	switch appConfig.LogConf.LogLevel {
	case "debug":
		level = xlog.XLogLevelDebug
	case "trace":
		level = xlog.XLogLevelTrace
	case "info":
		level = xlog.XLogLevelInfo
	case "warn":
		level = xlog.XLogLevelWarn
	case "error":
		level = xlog.XLogLevelError
	default:
		level = xlog.XLogLevelDebug
	}

	err = xlog.Init(logType, level, appConfig.LogConf.Filename, appConfig.LogConf.Module)
	return
}

func main() {
	err := initConfig("./conf/config.ini")
	if err != nil {
		panic(fmt.Sprintf("init config failed, err:%v", err))
	}

	err = initLog()
	if err != nil {
		panic(fmt.Sprintf("init logs failed, err:%v", err))
	}

	xlog.LogDebug("init log succ")

	address := strings.Split(appConfig.KafkaConf.Address, ",")
	err = kafka.Init(address, appConfig.KafkaConf.QueueSize)
	if err != nil {
		panic(fmt.Sprintf("init kafka client failed, err:%v", err))
	}

	xlog.LogDebug("init kafka succ")

	err = tailf.Init(appConfig.CollectLogConf.LogFilenames)
	if err != nil {
		panic(fmt.Sprintf("init tailf client failed, err:%v", err))
	}

	xlog.LogDebug("init tailf succ")
	err = run()
	if err != nil {
		xlog.LogError("run failed, err:%v", err)
		return
	}
	xlog.LogDebug("run finished")
}