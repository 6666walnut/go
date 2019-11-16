/**
* @Author: Chicken dishes
* @Date: 2019/10/22 9:28
 */
package main

import "fmt"

type person struct {
	name, gender string
}

// 拷贝
func f(x person) {
	x.gender = "nv1"
}

// 传指针
func f1(x *person) {
	x.gender = "nv2"
}

func main() {
	var p person
	p.name = "li"
	p.gender = "nan"
	f(p)
	fmt.Printf("this is f:%v \n", p)
	f1(&p)
	fmt.Printf("this is f1:%v \n", p)

	var p1 = new(person)
	p1.name = "tan"
	p1.gender = "baomi"
	fmt.Printf("type:%T value:%v\n", p1, &p1)

	var p2 = person{
		name:   "deng",
		gender: "nv",
	}
	fmt.Println(p2)

	p3 := person{
		"name",
		"gender",
	}
	fmt.Println(p3)
}
