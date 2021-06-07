/**
* @Author: Chicken dishes
* @Date: 2021/6/5 17:06
**/

package main

/**
 * 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
 * 如果有多对数字的和等于s，则输出任意一对即可。
 */
func twoSum(nums []int, target int) []int {
	var numMap = make(map[int]int, 0)
	for k, v := range nums {
		_, ok := numMap[target-v]
		if ok {
			return []int{target - v, v}
		}
		numMap[v] = k
	}
	return nil
}

func main() {
	twoSum([]int{3, 2, 4}, 6)
}
