/**
* @Author: Chicken dishes
* @Date: 2019/10/20 21:34
 */

package main

import (
	"fmt"
)

func main() {

	// 匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	//
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
