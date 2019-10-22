/**
* @Author: Chicken dishes
* @Date: 2019/10/22 8:33
 */
package main

import "fmt"

func f1(name string) {
	fmt.Println("hello", name)
}

func f2(f func(string), name string) {
	f(name)
}

func f4(x, y int) int {
	return x + y
}

func f3() func(int, int) int {
	return f4
}

// 闭包

func f5(f func(string), name string) func() {
	fmt.Println("this is f5")
	return func() {
		f(name)
	}
}

func f6(f func())  {
	fmt.Println("this is f6")
	f()
}

func main() {
	f2(f1, "NB")
	ret := f3()
	fmt.Printf("%T \n", ret)
	myNum := ret(10, 20)
	fmt.Println(myNum)

	fc := f5(f1,"tan")
	f6(fc)

}
