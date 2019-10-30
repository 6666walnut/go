/**
* @Author: Chicken dishes
* @Date: 2019/10/30 8:49
 */
package _8mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string //日志路径
	fileName    string //日志名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}

// 初始化文件
func (f *FileLogger) initFile() (error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) Close() (error) {
	err := f.fileObj.Close()
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	err = f.errFileObj.Close()
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	return nil
}

// 判断日志等级

func (f *FileLogger) enable(level LogLevel) bool {
	return f.Level <= level
}

// 文件超过大小进行重命名
func (f *FileLogger) checkSize(file *os.File) (*os.File,error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return file,err
	}
	if fileInfo.Size() > f.maxFileSize {
		// 切割文件大小
		// 1.关闭当前文件
		file.Close()
		// 2.备份rename
		logName := path.Join(f.filePath, fileInfo.Name())
		nowStr := time.Now().Format("20060102150405000")
		newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
		err = os.Rename(logName, newLogName)
		if err != nil {
			panic("rename is err")
			return file,err
		}
		// 3. 打开新的
		fileObj, err := oos.OpenFile(logName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("open file is failed")
			return file,err
		}
		return fileObj,nil
	}
	return file,err
}

// 写日志
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNu := getInfo(3)

		_, err := fmt.Fprintf(f.fileObj, "%s|[%s] %s |%s| %d| %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNu, msg)
		if err != nil {
			panic("this is err for write to ")
		}
		if lv >= ERROR {
			_, err = fmt.Fprintf(f.errFileObj, "%s|[%s] %s |%s| %d| %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNu, msg)
			if err != nil {
				panic("this is err for write to error log")
			}
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Warn(format string, a ...interface{}) {
	f.log(WARN, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)

}
