package main


import (
	"github.com/gostudy03/xlog"
	"flag"
	"fmt"
)

func logic() {
	for {
		xlog.LogDebug("dskskdskfksdf, user_id:%d username:%s", 388338, "username")
		xlog.LogTrace("dskskdskfksdf")
		xlog.LogInfo("dskskdskfksdf")
		xlog.LogWarn("dskskdskfksdf")
		xlog.LogError("dskskdskfksdf")
	}
}

/*
func testGetLine() {
	filename, funcName , lineNo  := xlog.GetLineInfo(2)
	fmt.Printf("filename:%s funcName:%s line:%d\n", filename, funcName, lineNo)
}*/

func main() {

	//testGetLine()

	var logTypeStr string
	flag.StringVar(&logTypeStr, "type", "console",  "please input logger type")
	flag.Parse()

	var logType int
	if (logTypeStr == "file") {
		logType = xlog.XLogTypeFile
	} else {
		logType = xlog.XLogTypeConsole
		
	}

	xlog.LogDebug("log type is %v", logType)

	_ = logType
	err := xlog.Init(logType, xlog.XLogLevelDebug, "C:/tmp/xlog.log", "xlog_example")
	if err != nil {
		fmt.Printf("logger init failed\n")
		return
	}
	logic()
	xlog.Close()
	fmt.Printf("close return")
}