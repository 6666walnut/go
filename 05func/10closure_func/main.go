/**
* @Author: Chicken dishes
* @Date: 2019/10/20 22:12
 */

package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {

	jpgFunc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFunc("test"))

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
}
