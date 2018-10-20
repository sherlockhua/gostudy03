package main

import (
	"strings"
	"github.com/gostudy03/oconfig"
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/logagent/kafka"
	"github.com/gostudy03/logagent/tailf"
	"github.com/gostudy03/logagent/etcd"
	"github.com/gostudy03/logagent/common"
	"fmt"
)

var (
	appConfig common.AppConfig
)

func initConfig(filename string) (err error) {

	err = oconfig.UnMarshalFile(filename, &appConfig)
	if err != nil {
		return
	}

	xlog.LogDebug("read config succ, config:%#v", appConfig)
	return
}

func run() (err error){

	//不断检测etcd配置是否有变更，如果有变更，那么需要对日志收集任务进行管理。
	tailf.Run()
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

	//初始化etcd client
	address = strings.Split(appConfig.EtcdConf.Address, ",")
	err = etcd.Init(address, appConfig.EtcdConf.EtcdKey)
	if err != nil {
		panic(fmt.Sprintf("init etcd client failed, err:%v", err))
	}
	xlog.LogDebug("init etcd succ, address:%v", address)

	logCollectConf, err := etcd.GetConfig(appConfig.EtcdConf.EtcdKey)
	xlog.LogDebug("etcd conf:%#v", logCollectConf)

	watchCh := etcd.Watch()
	err = tailf.Init(logCollectConf, watchCh)
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