/**
* @Author: Chicken dishes
* @Date: 2019/10/26 16:21
 */

package main

import "fmt"

// 空接口：没有必要起名字，任意的变量都能保存

func show(a interface{})  {
	fmt.Printf("type:%T,%v",a,a)
}
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "t1"
	m1["age"] = 18
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)
	show("tan")
	show(nil)
	show([...]string{"唱", "跳", "rap"})
}
