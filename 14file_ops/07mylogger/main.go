/**
* @Author: Chicken dishes
* @Date: 2019/10/29 8:47
 */

package main

import (
	"goStart/14file_ops/08mylogger"
	"time"
)

func main() {
	log := _8mylogger.NewFileLogger("info", "./", "tan.log",10*1024*1024)
	for {
		log.Debug("this is debug,id:%d", 1)
		log.Info("this is info,id:%d", 1)
		log.Warn("this is warn")
		log.Error("this is error")
		time.Sleep(5 * time.Second)
	}

}
