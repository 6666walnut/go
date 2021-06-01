/**
* @Author: Chicken dishes
* @Date: 2021/5/30 10:05
**/

package main

import (
	"math"
)

/**
 * 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 */

func majorityElement(nums []int) int {

	var numsMap = make(map[int]int)
	numsLen := int(math.Ceil(float64(len(nums)) / 2))
	for _, k := range nums {
		numsMap[k]++
		if numsMap[k] >= numsLen {
			return k
		}
	}
	return 0
}

func main() {
	majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2})
}
