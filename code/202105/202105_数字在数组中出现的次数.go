/**
* @Author: Chicken dishes
* @Date: 2021/5/31 21:15
**/

package main

import "fmt"

/**
 * 统计一个数字在排序数组中出现的次数。
 */

func search(nums []int, target int) int {

	mid := len(nums) / 2
	l, r := 0, len(nums)-1
	for r >= l {
		mid = (l + r) / 2
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			num := 0
			for i := mid; i >= 0 && target == nums[i]; i-- {
				num++
			}
			for j := mid + 1; j < len(nums)-1 && target == nums[j]; j++ {
				num++
			}
			return num
		}
	}

	return 0
}

func main() {
	c := search([]int{5}, 5)
	fmt.Println(c)
}
