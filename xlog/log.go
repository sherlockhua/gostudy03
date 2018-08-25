package xlog

type XLog interface {
	Init() error
	LogDebug(fmt string, args ...interface{})
	LogTrace(fmt string, args ...interface{})
	LogInfo(fmt string, args ...interface{})
	LogWarn(fmt string, args ...interface{})
	LogError(fmt string, args ...interface{})
	LogFatal(fmt string, args ...interface{})

	Close()
	SetLevel(level int)
}

func NewXLog(logType, level int, filename, module string) XLog {

	var logger XLog
	switch logType {
	case XLogTypeFile:
		logger = NewXFile(level, filename, module)
	case XLogTypeConsole:
		logger = NewXConsole(level, module)
	default:
		logger = NewXFile(level, filename, module)
	}
	return logger
}
