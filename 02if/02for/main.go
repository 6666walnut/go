/**
* @Author: Chicken dishes
* @Date: 2019/10/17 8:20
 */

package main

import "fmt"

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种2
	var j = 5
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// 无限循环
	//for {
	//	fmt.Println("123")
	//}

	// fo range 循环

	s := "hello中国"
	for i, v := range s {
		fmt.Println(i, v)
	}
}
