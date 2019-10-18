/**
* @Author: Chicken dishes
* @Date: 2019/10/18 7:37
 */

package main

import "fmt"


// 找出数组和为指定的两个元素的下标,比如从数组[1,3,5,7,8]中找出和位8的2个元素(0,3)和(1,2)
func main() {

	arr1 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i]+arr1[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
