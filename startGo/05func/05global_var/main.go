/**
* @Author: Chicken dishes
* @Date: 2019/10/20 21:19
 */

package main

import "fmt"

var x = 100

func f1() {
	// 函数中的查找变量的顺序
	// 1.先在函数内部查找
	// 2.找不到就往函数外部查找，一直找到全局
	fmt.Println(x)
}

func main() {
	f1()

	// 语句块作用域
	if i := 10; i < 18 {
		fmt.Println(i)
	}

}
