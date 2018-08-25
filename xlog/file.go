package xlog

import (
	"os"
)

type XFile struct {
	filename string
	file *os.File
	*XLogBase
}

func NewXFile(level int, filename, module string) XLog {
	logger := &XFile{
		filename: filename,
	}

	logger.XLogBase = &XLogBase{
		level : level,
		module : module,
	}
	return logger
}

func (c *XFile) Init() (err error) {
	c.file, err = os.OpenFile(c.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	return
}

func (c *XFile) LogDebug(format string, args ...interface{}) {
	if c.level > XLogLevelDebug {
		return
	}


	logData := c.formatLogger(XLogLevelDebug, c.module, format, args...)
	c.writeLog(c.file, logData)
}

func (c *XFile) LogTrace(format string, args ...interface{}) {
	if c.level > XLogLevelTrace {
		return
	}


	logData := c.formatLogger(XLogLevelTrace, c.module, format, args...)
	c.writeLog(c.file, logData)
}

func (c *XFile) LogInfo(format string, args ...interface{}) {
	if c.level > XLogLevelInfo {
		return
	}


	logData := c.formatLogger(XLogLevelInfo, c.module, format, args...)
	c.writeLog(c.file, logData)
}

func (c *XFile) LogWarn(format string, args ...interface{}) {
	if c.level > XLogLevelWarn {
		return
	}


	logData := c.formatLogger(XLogLevelWarn, c.module, format, args...)
	c.writeLog(c.file, logData)
}

func (c *XFile) LogError(format string, args ...interface{}) {
	if c.level > XLogLevelError {
		return
	}


	logData := c.formatLogger(XLogLevelError, c.module, format, args...)
	c.writeLog(c.file, logData)
}

func (c *XFile) LogFatal(format string, args ...interface{}) {
	if c.level > XLogLevelFatal {
		return
	}

	logData := c.formatLogger(XLogLevelFatal, c.module, format, args...)
	c.writeLog(c.file, logData)
}

func (c *XFile) SetLevel(level int) {
	c.level = level
}

func (c*XFile) Close() {
	if c.file != nil {
		c.file.Close()
	}
}