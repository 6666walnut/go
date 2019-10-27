/**
* @Author: Chicken dishes
* @Date: 2019/10/26 16:18
 */

package main

import "fmt"

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接受接口的所有方法
func (c *cat) move() {
	fmt.Println("猫动！")
}
func (c *cat) eat(s string) {
	fmt.Printf("猫吃！%s\n", s)
}

func main() {
	var a1 animal
	c1 := cat{
		"tom",
		4,
	}
	a1 = &c1
	fmt.Println(a1)
}
