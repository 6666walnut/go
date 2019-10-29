/**
* @Author: Chicken dishes
* @Date: 2019/10/29 8:13
 */

package main

import (
	"fmt"
	"time"
)

func f1() {

}

func f2() {
	now := time.Now()
	fmt.Println(now)
	// 获取明天的时间
	timeNow, err := time.Parse("2016-01-02 15:04:05", "2019-10-30 08:18:38")
	fmt.Println(timeNow)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}
	timeObj, err := time.ParseInLocation("2016-01-02 15:04:05", "2019-10-30 08:18:38", loc)
	fmt.Println(timeObj)
	fmt.Println(now.Sub(timeNow))
	fmt.Println(now.Sub(timeObj))
}

func main() {
	f2()
}
