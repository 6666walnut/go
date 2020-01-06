/**
* @Author: Chicken dishes
* @Date: 2020/1/6 22:16
 */
package main

import "fmt"

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i] + nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	myret := maxSubArray(nums)
	fmt.Println(myret)
}
