/**
* @Author: Chicken dishes
* @Date: 2019/10/31 7:20
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a *int
	// 判断有没有值
	fmt.Println("var a *int isnil", reflect.ValueOf(a).IsNil())
	b:= struct {

	}{}
	fmt.Println("不存在",reflect.ValueOf(b).FieldByName("a").IsValid())
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}
