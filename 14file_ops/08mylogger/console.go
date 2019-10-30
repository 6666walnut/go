/**
* @Author: Chicken dishes
* @Date: 2019/10/29 8:44
 */

package _8mylogger

import (
	"fmt"
	"time"
)

type LogLevel uint16

type ConsoleLogger struct {
	Level LogLevel
}

// 函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		level,
	}
}

// 判断日志等级

func (c ConsoleLogger) enable(level LogLevel) bool {
	return c.Level <= level
}

// 写日志
func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNu := getInfo(3)
		levelStr := getLogString(lv)
		fmt.Printf("%s|[%s] %s |%s| %d| %s\n", now.Format("2006-01-02 15:04:05"), levelStr, funcName, fileName, lineNu, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Warn(format string, a ...interface{}) {
	c.log(WARN, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
