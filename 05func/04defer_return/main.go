/**
* @Author: Chicken dishes
* @Date: 2019/10/20 21:13
 */

package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func f5() (x int) {
	defer func(int) int {
		x++
		return x
	}(x)
	return 5
}

func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5
}

// defer

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}


// a := 1
// b := 2
// defer calc("1", 1, calc("10", 1, 2))  1 1 3 4
// calc("10", 1, 2)   10 1 2 3
// a = 0
// defer calc("2", 0, calc("20", 0, 2))  2 0 2 2
// calc("20", 0, 2)   20 0 2 2
// b = 1

func f7() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}





func main() {
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
	//fmt.Println(f5())
	//fmt.Println(f6())
	f7()
}
