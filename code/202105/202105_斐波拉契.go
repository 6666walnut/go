/**
* @Author: Chicken dishes
* @Date: 2021/5/24 20:27
**/

package main

import "fmt"

func fib(n int) int {

	if n == 0 {
		return 0
	}
	// 迭代
	var x, y = 0, 1

	for i := 0; i < n-1; i++ {
		y, x = (x+y)%1000000007, y%1000000007
	}
	return y
}

func main() {
	y := fib(45)
	fmt.Println(y)
}
