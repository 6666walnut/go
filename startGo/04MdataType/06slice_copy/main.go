/**
* @Author: Chicken dishes
* @Date: 2019/10/18 8:50
 */

package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 将a1中的索引位1,3的元素删除
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)

	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	s1 = append(s1[:1],s1[2:]...)//修改了底层数组
	fmt.Println(s1,x1)
}
