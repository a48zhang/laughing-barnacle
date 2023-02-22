package service

import (
	"fmt"
	"path"
	"runtime"
)

func ErrorSender() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func Debugger(skip int) (fileName string, funcName string, line int) {
	pc, file, line, _ := runtime.Caller(skip)
	fmt.Println(pc, file)
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return fileName, funcName, line
}
