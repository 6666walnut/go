/**
* @Author: Chicken dishes
* @Date: 2019/10/20 21:23
 */

package main

import (
	"fmt"
)

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(a, b int) int {
	return a + b
}

func f5(x func() int) func(int, int) int {
	return f4
}

func main() {
	a := f1
	fmt.Printf("%T \n", a)
	b := f2
	fmt.Printf("%T \n", b)

	f3(f2)

	fmt.Println(f5(f2))
}
