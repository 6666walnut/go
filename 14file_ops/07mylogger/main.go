/**
* @Author: Chicken dishes
* @Date: 2019/10/29 8:47
 */

package main

import (
	"goStart/14file_ops/mylogger"
	"time"
)

func main() {
	log := mylogger.NewLog("Info")
	for {
		log.Debug("this is debug")
		log.Info("this is info")
		log.Warn("this is warn")
		log.Error("this is error")
		time.Sleep(5 * time.Second)
	}

}
