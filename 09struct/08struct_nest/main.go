/**
* @Author: Chicken dishes
* @Date: 2019/10/24 8:45
 */
package main

import "fmt"

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	addr address
}

type company struct {
	name string
	address
}

func main() {
	p1 := person{
		name: "p1",
		age:  18,
		addr: address{
			province: "湖北",
			city:     "武汉",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.addr.city)

	c1 := company{
		name: "c1",
		address: address{
			province: "湖北",
			city:     "武汉",
		},
	}
	fmt.Println(c1.city)
}
