/**
* @Author: Chicken dishes
* @Date: 2019/10/26 15:59
 */

package main

import "fmt"

type animal interface {
	move()
	eat(someting string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫动！")
}
func (c cat) eat(s string) {
	fmt.Printf("猫吃！%s\n", s)
}

type chiceken struct {
	feet int8
}

func (c chiceken) move() {
	fmt.Println("激动！")
}
func (c chiceken) eat(string2 string) {
	fmt.Println("吃鸡！")
}

func main() {
	var a1 animal
	bc := cat{
		name: "小",
		feet: 4,
	}
	a1 = bc
	fmt.Printf("%T\n",a1)
	a1.eat("小黄鱼")

	kfc := chiceken{
		feet: 2,
	}
	a1 = kfc
	fmt.Printf("%T\n",a1)
}
