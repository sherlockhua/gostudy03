package xlog

import (
	"fmt"
	"path/filepath"
	"time"
)

type XConsole struct {
	level  int
	module string
}

func NewXConsole(level int, module string) XLog {
	logger := &XConsole{
		level:  level,
		module: module,
	}
	return logger
}

type LogData struct {
	timeStr  string
	levelStr string
	module   string
	filename string
	funcName string
	lineNo   int
	data     string
}

func formatLogger(level int, module string, format string, args ...interface{}) *LogData {

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

func (c *XConsole) LogDebug(format string, args ...interface{}) {

	if c.level > XLogLevelDebug {
		return
	}

	logData := formatLogger(XLogLevelDebug, c.module, format, args...)
	fmt.Printf("%s %s %s (%s:%s:%d) %s\n",
		logData.timeStr, logData.levelStr, logData.module, logData.filename,
		logData.funcName, logData.lineNo, logData.data)
}

func (c *XConsole) LogTrace(format string, args ...interface{}) {
	if c.level > XLogLevelTrace {
		return
	}

	logData := formatLogger(XLogLevelTrace, c.module, format, args...)
	fmt.Printf("%s %s %s (%s:%s:%d) %s\n",
		logData.timeStr, logData.levelStr, logData.module, logData.filename,
		logData.funcName, logData.lineNo, logData.data)
}

func (c *XConsole) LogInfo(format string, args ...interface{}) {
	if c.level > XLogLevelInfo {
		return
	}

	logData := formatLogger(XLogLevelInfo, c.module, format, args...)
	fmt.Printf("%s %s %s (%s:%s:%d) %s\n",
		logData.timeStr, logData.levelStr, logData.module, logData.filename,
		logData.funcName, logData.lineNo, logData.data)
}

func (c *XConsole) LogWarn(format string, args ...interface{}) {
	if c.level > XLogLevelWarn {
		return
	}

	logData := formatLogger(XLogLevelWarn, c.module, format, args...)
	fmt.Printf("%s %s %s (%s:%s:%d) %s\n",
		logData.timeStr, logData.levelStr, logData.module, logData.filename,
		logData.funcName, logData.lineNo, logData.data)
}

func (c *XConsole) LogError(format string, args ...interface{}) {
	if c.level > XLogLevelError {
		return
	}

	logData := formatLogger(XLogLevelError, c.module, format, args...)
	fmt.Printf("%s %s %s (%s:%s:%d) %s\n",
		logData.timeStr, logData.levelStr, logData.module, logData.filename,
		logData.funcName, logData.lineNo, logData.data)
}

func (c *XConsole) LogFatal(format string, args ...interface{}) {
	if c.level > XLogLevelFatal {
		return
	}

	logData := formatLogger(XLogLevelFatal, c.module, format, args...)
	fmt.Printf("%s %s %s (%s:%s:%d) %s\n",
		logData.timeStr, logData.levelStr, logData.module, logData.filename,
		logData.funcName, logData.lineNo, logData.data)
}

func (c *XConsole) SetLevel(level int) {
	c.level = level
}
