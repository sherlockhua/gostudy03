package xlog

import (
	"runtime"
)

func getLineInfo(skip int) (filename, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		fun := runtime.FuncForPC(pc)
		funcName = fun.Name()
	}

	filename = file
	lineNo = line
	return
}
