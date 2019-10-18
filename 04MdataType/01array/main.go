/**
* @Author: Chicken dishes
* @Date: 2019/10/18 7:16
 */

package main

import "fmt"

// 数组

// 存放元素的容器
// 必须知道存放的元素的类型和容器(长度)
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("%T", a1)
	fmt.Printf("%T", a2)

	// 数组初始化
	// 如果不出初始化：默认都是零值
	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 2.初始化方式2
	a3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a3)

	// 初始化方式3
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)

	// 索引遍历
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
	// for range 遍历
	for i, v := range a4 {
		fmt.Println(i, v)
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

}
