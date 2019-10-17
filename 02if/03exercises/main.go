/**
* @Author: Chicken dishes
* @Date: 2019/10/17 8:27
 */

package main

import "fmt"

// 99乘法表
func main() {
	for i := 1; i < 10; i++ {
		for j := 1;j <= i; j++{
			fmt.Printf("%d x %d = %d \t",i,j,i*j)
		}
		fmt.Printf("\n")
	}
}
