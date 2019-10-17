/**
* @Author: Chicken dishes
* @Date: 2019/10/17 8:45
 */

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			break //跳出for循环
		}
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //跳出for循环
		}
		fmt.Println(i)
	}
}
