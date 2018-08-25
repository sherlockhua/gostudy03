package xlog

import (
	"fmt"
)

type XFile struct {
	level    int
	filename string
	module   string
}

func NewXFile(level int, filename, module string) XLog {
	logger := &XFile{
		level:    level,
		filename: filename,
		module:   module,
	}
	return logger
}

func (c *XFile) LogDebug(format string, args ...interface{}) {
	fmt.Printf("log debug of file\n")
}

func (c *XFile) LogTrace(format string, args ...interface{}) {
	fmt.Printf("log trace of file\n")
}

func (c *XFile) LogInfo(format string, args ...interface{}) {
	fmt.Printf("log info of file\n")
}

func (c *XFile) LogWarn(format string, args ...interface{}) {
	fmt.Printf("log warn of file\n")
}

func (c *XFile) LogError(format string, args ...interface{}) {
	fmt.Printf("log error of file\n")
}

func (c *XFile) LogFatal(format string, args ...interface{}) {
	fmt.Printf("log fatal of file\n")
}

func (c *XFile) SetLevel(level int) {
	c.level = level
}
