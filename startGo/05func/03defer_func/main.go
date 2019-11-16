/**
* @Author: Chicken dishes
* @Date: 2019/10/20 21:09
 */

package main

import "fmt"

func deferDemo()  {
	defer  fmt.Println("呵呵呵呵")
	fmt.Println("start")
	defer  fmt.Println("哈哈哈哈哈")
	fmt.Println("end")
}

func main()  {
	deferDemo()
}