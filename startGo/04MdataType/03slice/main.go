/**
* @Author: Chicken dishes
* @Date: 2019/10/18 7:46
 */

package main

import "fmt"

// 切片 slice
func main() {
	// 切片定义
	var s1 []int    //定义一个存放int 类型元素的切片
	var s2 []string //定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"wuhan", "beijing", "shanghai"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 长度和容量

	fmt.Printf("len(s1):%d  cap(s1):%d", len(s1), cap(s1))
	fmt.Printf("len(s2):%d  cap(s2):%d \n", len(s2), cap(s2))

	// 2.由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := a1[0:4]
	fmt.Println(s3)

	s4 := a1[1:6]
	fmt.Println(s4)
	fmt.Printf("len(s1):%d  cap(s1):%d \n", len(s3), cap(s3))
	fmt.Printf("len(s2):%d  cap(s2):%d \n", len(s4), cap(s4))

	// 切片是引用
	a1[1] = 10
	fmt.Println(s4)
	s4[2] = 20

	fmt.Println(a1)
}
