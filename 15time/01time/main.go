/**
* @Author: Chicken dishes
* @Date: 2019/10/27 14:09
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())

	ret := time.Unix(1572156690, 0)
	fmt.Println(ret)

	// 时间间隔
	fmt.Println(time.Second)
	// now + 24 小时
	fmt.Println(now.Add(24 * time.Hour))
	// 定时器
	//timer := time.Tick(time.Second)
	//{
	//	for t := range timer {
	//		fmt.Println(t)
	//	}
	//}

	//格式化时间 把语言中的时间对象，转换成字符串类型的时间
	// 2019-10-27
	fmt.Println(now.Format("2006-01-01"))
	fmt.Println(now.Sub(now.Add(24 * time.Hour)))
	// 按照对呀的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2019-10-27")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())



	now = now.UTC()
	fmt.Printf("现在时间:%v",now)

	// sleep
	n := 5
	fmt.Println("开始sleep了")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5秒钟过去了")
}
