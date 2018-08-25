package main


import (
	"github.com/gostudy03/xlog"
	"flag"
	//"fmt"
)

func logic(logger xlog.XLog) {
	logger.LogDebug("dskskdskfksdf, user_id:%d username:%s", 388338, "username")
	logger.LogTrace("dskskdskfksdf")
	logger.LogInfo("dskskdskfksdf")
	logger.LogWarn("dskskdskfksdf")
	logger.LogError("dskskdskfksdf")
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

	logger := xlog.NewXLog(logType, xlog.XLogLevelDebug, "", "xlog_example")
	logic(logger)
}