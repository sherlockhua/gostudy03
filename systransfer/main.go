package main

import (
	"github.com/gostudy03/oconfig"
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/systransfer/common"
	"github.com/gostudy03/systransfer/kafka"
	"github.com/gostudy03/systransfer/influx"
)


func initLog() (err error) {
	
		var logType int
		var level int
	
		if common.AppConf.LogConf.LogType == "console" {
			logType = xlog.XLogTypeConsole
		} else {
			logType = xlog.XLogTypeFile
		}
	
		switch common.AppConf.LogConf.LogLevel {
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
	
		err = xlog.Init(logType, level, common.AppConf.LogConf.Filename, common.AppConf.LogConf.Module)
		return
	}

func main() {
	err := oconfig.UnMarshalFile("./conf/systransfer.ini", &common.AppConf)
	if err != nil {
		xlog.LogError("unmarshal file failed, err:%v\n", err)
		return
	}

	xlog.LogDebug("app conf:%#v\n", common.AppConf)
	err = initLog() 
	if err != nil {
		xlog.LogError("init log failed, err:%v", err)
		return
	}

	err = influx.Init(common.AppConf.InfluxConf.Addr)
	if err != nil {
		xlog.LogError("init influx db failed, err:%v", err)
		return
	}
	
	err = kafka.Init(common.AppConf.KafkaConf.Addr, common.AppConf.KafkaConf.Topic)
	if err != nil {
		xlog.LogError("init kafka failed, err:%v", err)
		return
	}
	


	select {}
}