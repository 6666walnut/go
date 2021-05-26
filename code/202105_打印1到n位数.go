/**
* @Author: Chicken dishes
* @Date: 2021/5/25 20:04
**/

package main

import (
	"fmt"
	"math"
)

/**
 * 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
 */

func printNumbers(n int) []int {

	if n <= 0 {
		return nil
	}
	max := math.Pow(10, float64(n))
	var nums []int
	for i := 1; i < int(max); i++ {
		nums = append(nums, i)
	}
	return nums
}

func main() {
	nums := printNumbers(2)
	fmt.Println(nums)
}
