/**
* @Author: Chicken dishes
* @Date: 2020/5/19 23:07
 */

package main

import "fmt"

// 二分法
func mySqrt(x int) int {
	l, r, ret := 0, x, 0
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}

func main() {
	myInt := 10
	ret := mySqrt(myInt)
	fmt.Println(ret)
}
