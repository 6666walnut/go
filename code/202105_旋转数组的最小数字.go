/**
* @Author: Chicken dishes
* @Date: 2021/5/24 20:48
**/

package main

import "fmt"

/**
 * 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，
 * 输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
 */

func minArray(numbers []int) int {
	numbersLen := len(numbers)
	var index int
	for i := 0; i < numbersLen-1; i++ {
		if numbers[i] <= numbers[i+1] {
			continue
		}
		index = i + 1
		break
	}
	return numbers[index]
}

func main() {
	numbers := []int{3, 3, 1}
	fmt.Println(minArray(numbers))
}
