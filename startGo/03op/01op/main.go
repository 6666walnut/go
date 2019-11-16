/**
* @Author: Chicken dishes
* @Date: 2019/10/17 9:01
 */

package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)
	// 算术运算
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	b--
	fmt.Println(a)
	fmt.Println(b)

	//关系运算
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	// 逻辑运算符

	age := 22
	if age > 18 && age < 60 {
		fmt.Println("上班")
	} else {
		fmt.Println("no work")
	}

	if age > 18 || age < 60 {
		fmt.Println("上班")
	} else {
		fmt.Println("no work")
	}

	isMarried := false

	fmt.Println(!isMarried)

	// 位运算：针对2进制
	// 5的二进制 101
	// 2的二进制 10

	// &: 按位于（2位均为1）
	fmt.Println(5 & 2)
	// |: 按位或 (2位一样为1)
	fmt.Println(5 | 2)
	// ^: 按位异 (2位不一样位1）
	fmt.Println(5 ^ 2)

	// <<:将二进制位左移指定位数
	fmt.Println(5 << 1)  //1010 => 10
	fmt.Println(1 << 10) // 10000000000 =>1024
	// >>:将二进制位右移指定位数
	fmt.Println(5 >> 3) // 0=>0
}
