/**
* @Author: Chicken dishes
* @Date: 2020/1/2 21:04
 */
package main

import (
	"fmt"
)

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//你可以假设数组中无重复元素。
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
