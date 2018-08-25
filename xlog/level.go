package xlog

const (
	XLogLevelDebug = iota
	XLogLevelTrace
	XLogLevelInfo
	XLogLevelWarn
	XLogLevelError
	XLogLevelFatal
)

const (
	XLogTypeFile = iota
	XLogTypeConsole
)

func getLevelStr(level int) string {
	switch level {
	case XLogLevelDebug:
		return "DEBUG"
	case XLogLevelTrace:
		return "TRACE"
	case XLogLevelInfo:
		return "INFO"
	case XLogLevelWarn:
		return "WARN"
	case XLogLevelError:
		return "ERROR"
	case XLogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
