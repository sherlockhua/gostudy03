package xlog

import (
	"os"
)

type XConsole struct {
	*XLogBase
}

func NewXConsole(level int, module string) XLog {

	logger := &XConsole{}
	logger.XLogBase = &XLogBase{
		level : level,
		module : module,
	}

	return logger
}

func (c *XConsole) Init() error {
	return nil
}

func (c *XConsole) LogDebug(format string, args ...interface{}) {

	if c.level > XLogLevelDebug {
		return
	}

	logData := c.formatLogger(XLogLevelDebug, c.module, format, args...)
	c.writeLog(os.Stdout, logData)
}

func (c *XConsole) LogTrace(format string, args ...interface{}) {
	if c.level > XLogLevelTrace {
		return
	}

	logData := c.formatLogger(XLogLevelTrace, c.module, format, args...)
	c.writeLog(os.Stdout, logData)
}

func (c *XConsole) LogInfo(format string, args ...interface{}) {
	if c.level > XLogLevelInfo {
		return
	}

	logData := c.formatLogger(XLogLevelInfo, c.module, format, args...)
	c.writeLog(os.Stdout, logData)
}

func (c *XConsole) LogWarn(format string, args ...interface{}) {
	if c.level > XLogLevelWarn {
		return
	}

	logData := c.formatLogger(XLogLevelWarn, c.module, format, args...)
	c.writeLog(os.Stdout, logData)
}

func (c *XConsole) LogError(format string, args ...interface{}) {
	if c.level > XLogLevelError {
		return
	}

	logData := c.formatLogger(XLogLevelError, c.module, format, args...)
	c.writeLog(os.Stdout, logData)
}

func (c *XConsole) LogFatal(format string, args ...interface{}) {
	if c.level > XLogLevelFatal {
		return
	}

	logData := c.formatLogger(XLogLevelFatal, c.module, format, args...)
	c.writeLog(os.Stdout, logData)
}

func (c *XConsole) SetLevel(level int) {
	c.level = level
}

func (c *XConsole) Close() {

}