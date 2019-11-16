/**
* @Author: Chicken dishes
* @Date: 2019/10/22 9:43
 */
package main

import "fmt"

type x struct {
	a int8
	b int8
	c int8
}

func main() {
	m := x{
		a: 100,
		b: 100,
		c: 100,
	}
	fmt.Printf("%p \n",&m.a)
	fmt.Printf("%p \n",&m.b)
	fmt.Printf("%p \n",&m.c)
}
