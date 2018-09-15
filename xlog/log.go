package xlog


var logger XLog = newXLog(XLogTypeConsole, XLogLevelDebug, "", "default")

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

func newXLog(logType, level int, filename, module string) XLog {

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

func Init(logType, level int, filename, module string) (error) {
	logger = newXLog(logType, level, filename, module)
	return logger.Init()
}

func LogDebug(fmt string, args ...interface{}) {
	 logger.LogDebug(fmt, args...)
}

func LogTrace(fmt string, args ...interface{}) {
	 logger.LogTrace(fmt, args...)
}

func LogInfo(fmt string, args ...interface{}) {
	 logger.LogInfo(fmt, args...)
}

func LogWarn(fmt string, args ...interface{}) {
	 logger.LogWarn(fmt, args...)
}

func LogError(fmt string, args ...interface{}) {
	 logger.LogError(fmt, args...)
}

func LogFatal(fmt string, args ...interface{}) {
	 logger.LogFatal(fmt, args...)
}

func Close() {
	logger.Close()
}

func SetLevel(level int) {
	logger.SetLevel(level)
}