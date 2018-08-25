package xlog

import (
	"os"
	"fmt"
	"time"
	"path/filepath"
)


type LogData struct {
	timeStr  string
	levelStr string
	module   string
	filename string
	funcName string
	lineNo   int
	data     string
}

type XLogBase struct  {
	level    int
	module   string
}

func (l *XLogBase) writeLog(file *os.File, logData *LogData) {

	fmt.Fprintf(file, "%s %s %s (%s:%s:%d) %s\n",
		logData.timeStr, logData.levelStr, logData.module, logData.filename,
		logData.funcName, logData.lineNo, logData.data)
}


func (l *XLogBase) formatLogger(level int, module string, format string, args ...interface{}) *LogData {
	
		now := time.Now()
		timeStr := now.Format("2006-01-02 15:04:05.000")
	
		levelStr := getLevelStr(level)
		filename, funcName, lineNo := getLineInfo(3)
	
		filename = filepath.Base(filename)
		data := fmt.Sprintf(format, args...)
	
		return &LogData{
			timeStr:  timeStr,
			levelStr: levelStr,
			module:   module,
			filename: filename,
			lineNo:   lineNo,
			funcName: funcName,
			data:     data,
		}
	}