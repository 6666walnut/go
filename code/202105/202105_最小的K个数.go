/**
* @Author: Chicken dishes
* @Date: 2021/5/31 19:32
**/

package main

import "fmt"

/*
 * 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
 */

func getLeastNumbers(arr []int, k int) []int {

	return quit_sort(arr, 0, len(arr)-1)[:k]
}

// 快排主要逻辑
func quit_sort(arr []int, l, r int) []int {

	if r > l {
		mid := partition(arr, l, r)
		quit_sort(arr, l, mid-1)
		quit_sort(arr, mid+1, r)
	}
	return arr
}

// 切分左右
func partition(arr []int, l, r int) int {
	x := arr[r]
	i := l - 1

	// 遍历数组
	for j := l; j < r; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{1, 4, 3, 6, 5}
	cc := getLeastNumbers(arr, 3)
	fmt.Println(cc)
}
