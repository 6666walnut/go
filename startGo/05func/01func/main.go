/**
* @Author: Chicken dishes
* @Date: 2019/10/19 9:37
 */

package main

import "fmt"

func sum(x int, y int) (ret int) {
	return x + y
}

func f1(x int, y int) {
	fmt.Println(x + y)
}

func f2() {
	fmt.Println("f2")
}

func f3() int {
	return 3
}

// 命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 使用明明的返回值可以return 后的省略
}

func f5() (int, string) {
	return 1, "NB"
}

// 参数类型简写
func f6(x, y int) int {
	return x + y
}

// 可变长参数
// 当参数中连续多个参数的类型的时候，我们可以将非最后一个参数的类型省略
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	f7("1", 2)
	f7("1", 2, 3, 4)
}
