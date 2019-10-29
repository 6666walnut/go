/**
* @Author: Chicken dishes
* @Date: 2019/10/29 8:37
 */

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("出现错误:%v\n", err)
		return
	}
	for {
		_,err := fileObj.Write([]byte("这是一个日志"))
		if err != nil {
			fmt.Printf("出现错误:%v\n", err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}
