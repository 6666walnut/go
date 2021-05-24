/**
* @Author: Chicken dishes
* @Date: 2020/1/8 22:24
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

func addBinary(a string, b string) string {
	d := len(a) - len(b)
	c := ""
	for i := 1; i <= int(math.Abs(float64(d))); i++ {
		c = "0" + c
	}
	if d > 0 {
		b = c + b
	} else {
		a = c + a
	}
	r, p := "", 0
	for i := len(a) - 1; i >= 0; i-- {
		s := int(a[i]-'0') + int(b[i]-'0') + p
		r = strconv.Itoa(s%2) + r
		p = s / 2
	}
	if p == 1 {
		r = "1" + r
	}
	return r
}

func main() {
	a := "1"
	b := "111"
	myStr := addBinary(a, b)
	fmt.Println(myStr)
}
