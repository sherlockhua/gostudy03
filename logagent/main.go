package main

import (
	"time"
	"sync"
	"strings"
	"github.com/gostudy03/oconfig"
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/logagent/kafka"
	"github.com/gostudy03/logagent/tailf"
	"github.com/gostudy03/logagent/etcd"
	"github.com/gostudy03/logagent/common"
	"github.com/gostudy03/logagent/collect_sys_info"
	"fmt"
)

var (
	appConfig common.AppConfig
	waitGroup sync.WaitGroup
)

func initConfig(filename string) (err error) {

	err = oconfig.UnMarshalFile(filename, &appConfig)
	if err != nil {
		return
	}

	xlog.LogDebug("read config succ, config:%#v", appConfig)
	return
}

func run(collectSystemInfoConfig *common.CollectSystemInfoConfig) (err error){

	waitGroup.Add(2)
	go collect_sys_info.Run(&waitGroup, collectSystemInfoConfig.Interval, collectSystemInfoConfig.Topic)
	//不断检测etcd配置是否有变更，如果有变更，那么需要对日志收集任务进行管理。
	go tailf.Run(&waitGroup)
	
	waitGroup.Wait()
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

	ip, err := common.GetLocalIP()
	if err != nil {
		xlog.LogError("get local ip failed, err:%v", err)
		return
	}
	
	xlog.LogDebug("local ip succ, ip:%v", ip)
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
	etcdKey := fmt.Sprintf(appConfig.EtcdConf.EtcdKey, ip)
	xlog.LogDebug("etcd key is %v", etcdKey)

	address = strings.Split(appConfig.EtcdConf.Address, ",")
	err = etcd.Init(address, etcdKey)
	if err != nil {
		panic(fmt.Sprintf("init etcd client failed, err:%v", err))
	}
	xlog.LogDebug("init etcd succ, address:%v", address)

	logCollectConf, err := etcd.GetConfig(etcdKey)
	xlog.LogDebug("etcd conf:%#v", logCollectConf)

	etcdCollectSystemInfoKey := fmt.Sprintf(appConfig.EtcdConf.EtcdCollectSystemInfoKey, ip)
	collectSystemInfoConfig, err := etcd.GetCollectSystemInfoConfig(etcdCollectSystemInfoKey)
	if err != nil {
		collectSystemInfoConfig = &common.CollectSystemInfoConfig{}
		collectSystemInfoConfig.Topic = "collect_system_info"
		collectSystemInfoConfig.Interval = 5 * time.Second
		xlog.LogError("get collect system info config failed, use default conf:%#v", collectSystemInfoConfig)
	}

	watchCh := etcd.Watch()
	err = tailf.Init(logCollectConf, watchCh)
	if err != nil {
		panic(fmt.Sprintf("init tailf client failed, err:%v", err))
	}

	xlog.LogDebug("init tailf succ")
	err = run(collectSystemInfoConfig)
	if err != nil {
		xlog.LogError("run failed, err:%v", err)
		return
	}
	xlog.LogDebug("run finished")
}