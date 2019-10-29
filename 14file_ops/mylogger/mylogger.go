/**
* @Author: Chicken dishes
* @Date: 2019/10/29 9:04
 */

package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	default:
		return UNKNOWN, errors.New("无效日志级别")
	}
}

func getInfo(n int) (funcName, fileName string, lineNu int) {
	pc, file, lineNu, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("不OK\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
