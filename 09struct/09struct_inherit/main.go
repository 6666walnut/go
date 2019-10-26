/**
* @Author: Chicken dishes
* @Date: 2019/10/24 8:52
 */
package main

import "fmt"

// 给 animal 一个移动的方法
type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s在移动\n",a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%swang\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "d1",
		},
	}
	d1.move()
	d1.wang()
}
