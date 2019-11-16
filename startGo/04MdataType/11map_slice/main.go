/**
* @Author: Chicken dishes
* @Date: 2019/10/19 9:25
 */

package main

import "fmt"

func main() {
	// 元素为map的切片
	var s1 = make([]map[int]string, 1, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "NB"
	fmt.Println(s1)

	// 值为切片的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
