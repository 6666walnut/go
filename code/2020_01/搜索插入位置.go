/**
* @Author: Chicken dishes
* @Date: 2020/1/2 21:04
 */
package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = left + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 5, 6}
	target := 7
	ret := searchInsert(nums, target)
	fmt.Println(ret)
}
