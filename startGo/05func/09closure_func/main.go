/**
* @Author: Chicken dishes
* @Date: 2019/10/20 21:58
 */

package main

import "fmt"

// 闭包底层原理：
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找
func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder(200)
	ret2 := ret(200)
	fmt.Println(ret2)
}
