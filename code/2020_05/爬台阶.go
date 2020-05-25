/**
* @Author: Chicken dishes
* @Date: 2020/5/25 21:14
 */

package main

import "fmt"

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶

// 迭代写法
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	num1, num2 := 1, 1
	for i := 1; i < n; i++ {
		num3 := num1 + num2
		num1 = num2
		num2 = num3
	}
	return num2
}
func main() {
	num := climbStairs(10)
	fmt.Print(num)
}
