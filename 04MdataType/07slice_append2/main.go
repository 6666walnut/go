/**
* @Author: Chicken dishes
* @Date: 2019/10/18 9:03
 */

package main

import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	s1 := a1[:]

	//删除索引位3的
	s1 = append(s1[0:1],s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)
}
