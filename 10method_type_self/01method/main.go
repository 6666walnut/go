/**
* @Author: Chicken dishes
* @Date: 2019/10/23 8:52
 */
package main

import "fmt"


// 给自定义类型添加方法
// 不能给其他包里面的类型添加方法

type myInt int

func (m myInt) hello() {
	fmt.Println("this is int")
}
func main() {
	m := myInt(100)
	m.hello()
}
