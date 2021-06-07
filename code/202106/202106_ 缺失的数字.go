/**
* @Author: Chicken dishes
* @Date: 2021/6/5 15:42
**/

package main

import "fmt"

/**
 * 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
 * 在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
 */

func missingNumber(nums []int) int {

	l, r := 0, len(nums)-1

	for r >= l {
		mid := (r + l) / 2
		if nums[mid] == mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

func main() {
	c := missingNumber([]int{0})
	fmt.Println(c)
}
