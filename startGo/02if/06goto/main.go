/**
* @Author: Chicken dishes
* @Date: 2019/10/17 8:53
 */

package main

import "fmt"

func main() {
	var flag bool = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break
			}
			fmt.Printf("%d-%c\n", i, j)
		}
		if flag {
			break
		}
	}
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx
			}
			fmt.Printf("%d-%c\n", i, j)
		}
		if flag {
			break
		}
	}
xx:
	fmt.Println("6")
}
