package xlog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type XFile struct {
	filename string
	file *os.File
	*XLogBase

	logChan  chan *LogData 
	wg *sync.WaitGroup
	curHour int
}

func NewXFile(level int, filename, module string) XLog {
	logger := &XFile{
		filename: filename,
	}

	logger.XLogBase = &XLogBase{
		level : level,
		module : module,
	}

	logger.curHour = time.Now().Hour()
	logger.wg = &sync.WaitGroup{}
	logger.logChan = make(chan *LogData, 10000)
	logger.wg.Add(1)
	go logger.syncLog()
	return logger
}

func (c *XFile) Init() (err error) {
	c.file, err = os.OpenFile(c.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	return
}

func (c *XFile)syncLog() {

	for data := range c.logChan {
		c.splitLog()
		c.writeLog(c.file, data)
	}

	c.wg.Done()
}

func (c *XFile) splitLog() {
	now := time.Now()
	if now.Hour() == c.curHour {
		return
	}

	c.curHour = now.Hour()
	c.file.Sync()
	c.file.Close()

	newFilename := fmt.Sprintf("%s-%04d-%02d-%02d-%02d", c.filename,
		 now.Year(), now.Month(), now.Day(), now.Hour())
	os.Rename(c.filename, newFilename)
	c.Init()
}

func (c *XFile) writeToChan(level int, module string, format string, args ...interface{}) {
	logData := c.formatLogger(level, module, format, args...)
	select {
	case c.logChan <- logData:
	default:
	}
}

func (c *XFile) LogDebug(format string, args ...interface{}) {
	if c.level > XLogLevelDebug {
		return
	}

	c.writeToChan(XLogLevelDebug, c.module, format, args...)
	//c.writeLog(c.file, logData)
}

func (c *XFile) LogTrace(format string, args ...interface{}) {
	if c.level > XLogLevelTrace {
		return
	}

	c.writeToChan(XLogLevelTrace, c.module, format, args...)
	/*
	logData := c.formatLogger(XLogLevelTrace, c.module, format, args...)
	c.writeLog(c.file, logData)
	*/
}

func (c *XFile) LogInfo(format string, args ...interface{}) {
	if c.level > XLogLevelInfo {
		return
	}

	c.writeToChan(XLogLevelInfo, c.module, format, args...)
	/*
	logData := c.formatLogger(XLogLevelInfo, c.module, format, args...)
	c.writeLog(c.file, logData)
	*/
}

func (c *XFile) LogWarn(format string, args ...interface{}) {
	if c.level > XLogLevelWarn {
		return
	}

	c.writeToChan(XLogLevelWarn, c.module, format, args...)
	/*
	logData := c.formatLogger(XLogLevelWarn, c.module, format, args...)
	c.writeLog(c.file, logData)
	*/
}

func (c *XFile) LogError(format string, args ...interface{}) {
	if c.level > XLogLevelError {
		return
	}

	c.writeToChan(XLogLevelError, c.module, format, args...)
	/*
	logData := c.formatLogger(XLogLevelError, c.module, format, args...)
	c.writeLog(c.file, logData)
	*/
}

func (c *XFile) LogFatal(format string, args ...interface{}) {
	if c.level > XLogLevelFatal {
		return
	}

	c.writeToChan(XLogLevelFatal, c.module, format, args...)
	/*
	logData := c.formatLogger(XLogLevelFatal, c.module, format, args...)
	c.writeLog(c.file, logData)
	*/
}

func (c *XFile) SetLevel(level int) {
	c.level = level
}

func (c*XFile) Close() {
	
	if c.logChan != nil {
		close(c.logChan)
	}
	
	c.wg.Wait()

	if c.file != nil {
		c.file.Sync()
		c.file.Close()
	}
}