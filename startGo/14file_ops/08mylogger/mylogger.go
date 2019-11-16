/**
* @Author: Chicken dishes
* @Date: 2019/10/29 9:04
 */

package _8mylogger

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

// 根据字符串返回日志等级
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
	pc, file, lineNu, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("不OK\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}


// 根据等级获取日志的等级字符串
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	default:
		return "不存在"
	}
}