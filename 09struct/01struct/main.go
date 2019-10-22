/**
* @Author: Chicken dishes
* @Date: 2019/10/22 9:19
 */
package main

import "fmt"

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person
	p.name = "li"
	p.age = 18
	p.gender = "nan"
	p.hobby = []string{"1", "2", "3"}
	fmt.Printf("type:%T  value:%v \n", p, p)

	var p1 person
	p1.name = "li"
	p1.age = 18
	fmt.Printf("type:%T  value:%v \n", p1, p1)

	// 匿名结构图

	var s struct {
		name string
		age  int
	}
	s.name = "tan"
	s.age = 18
	fmt.Printf("type:%T  value:%v \n", s, s)
}
