/**
* @Author: Chicken dishes
* @Date: 2021/5/24 20:39
**/

package main

import "fmt"

/**
 * 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
 * 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
 */

func numWays(n int) int {
	// 迭代
	var x, y = 0, 1
	for i := 0; i < n; i++ {
		y, x = (x+y)%1000000007, y%1000000007
	}
	return y
}

func main() {
	fmt.Println(numWays(0))
}
