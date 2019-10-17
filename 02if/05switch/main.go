/**
* @Author: Chicken dishes
* @Date: 2019/10/17 8:49
 */

package main

import "fmt"

func mian() {
	n := 5
	if n == 1 {
		fmt.Printf("这是1")
	} else if n == 2 {
		fmt.Printf("这是2")
	} else if n == 3 {
		fmt.Printf("这是3")
	} else if n == 4 {
		fmt.Printf("这是4")
	} else if n == 5 {
		fmt.Printf("这是5")
	}

	switch n {
	case 1:
		fmt.Printf("这是1")
	case 2:
		fmt.Printf("这是2")
	case 3:
		fmt.Printf("这是3")
	case 4:
		fmt.Printf("这是4")
		fallthrough
	case 5:
		fmt.Printf("这是5")
	default:
		fmt.Printf("这是其他数")
	}
}
