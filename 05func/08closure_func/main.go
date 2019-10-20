/**
* @Author: Chicken dishes
* @Date: 2019/10/20 21:38
 */

package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(x, y int), a, b int) func() {
	tmp := func() {
		f(a, b)
	}
	return tmp
}

// 要求:
// f1(f2)
func f4(f func(int, int), x, y int) func() {
	tmp := func() {
		fmt.Println("this is tmp")
		f(x, y)
	}
	return tmp
}

func main() {
	//ret := f3(f2, 5, 6)
	//ret()

	f1(f4(f2,100, 200))
}
