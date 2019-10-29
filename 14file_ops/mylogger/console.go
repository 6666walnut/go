/**
* @Author: Chicken dishes
* @Date: 2019/10/29 8:44
 */

package mylogger

import (
	"fmt"
	"time"
)

type LogLevel uint16

type Logger struct {
	Level LogLevel
}

func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		level,
	}
}

// 判断日志等级

func (l Logger) enable(level LogLevel) bool {
	return l.Level <= level
}

// 写日志
func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNu := getInfo(2)
	levelStr := getLogString(lv)
	fmt.Printf("[%s] [%s] [%s %s %d]%s\n", now.Format("2006-01-02 15:04:05"), levelStr, funcName, fileName, lineNu, msg)
}

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

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}

}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

func (l Logger) Warn(msg string) {
	if l.enable(WARN) {
		log(WARN, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}
