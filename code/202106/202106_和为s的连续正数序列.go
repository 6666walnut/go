/**
* @Author: Chicken dishes
* @Date: 2021/6/5 17:18
**/

package main

import "fmt"

/**
 * 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
 * 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
 */

func findContinuousSequence(target int) [][]int {
	var l, r, sum = 1, 2, 3
	var arr = make([][]int, 0)

	for r > l {
		if sum == target {
			temp := make([]int, 0)
			for i := l; i <= r; i++ {
				temp = append(temp, i)
			}
			arr = append(arr, temp)
		}
		if sum < target {
			r += 1
			sum += r
		} else {
			sum -= l
			l += 1
		}
	}

	return arr
}

func main() {
	a := findContinuousSequence(9)
	fmt.Println(a)
}
