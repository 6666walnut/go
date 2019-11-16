/**
* @Author: Chicken dishes
* @Date: 2019/10/18 8:01
 */

package main

import "fmt"

func main() {
	// 切片的本质：定义一块连续的内存
	s1 := make([]int, 5, 10)
	fmt.Printf("%d len:%d cap:%d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("%d len:%d cap:%d \n", s2, len(s2), cap(s2))

	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 10
	fmt.Println(s3, s4)
}
