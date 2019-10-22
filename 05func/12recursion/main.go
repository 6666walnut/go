/**
* @Author: Chicken dishes
* @Date: 2019/10/22 8:59
 */
package main

import "fmt"

// 计算阶乘
func f1(n uint64) uint64 {
	if n <= 1 {
		return 1
	} else {
		return n * f1(n-1)
	}
}

// 斐波拉契数列
func f2(num int) (result int) {
	if num <= 2 {
		return num
	}
	return f2(num-1) + f2(num-2)
}


func main() {
	ret := f1(5)
	fmt.Printf("5的阶乘为:%d\n", ret)

	ret1 := f2(6)
	fmt.Printf("斐波拉契数列:%d \n", ret1)


}
