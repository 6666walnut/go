/**
* @Author: Chicken dishes
* @Date: 2021/5/23 15:47
**/

package main

// 暴力解法
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum1(nums []int, target int) []int {
	var twoMap = map[int]int{}

	for k, v := range nums {
		j, ok := twoMap[target-v]
		if ok {
			return []int{j, k}
		}
		twoMap[v] = k
	}
	return nil
}
func main() {
	twoSum1([]int{3, 2, 4}, 6)
}
