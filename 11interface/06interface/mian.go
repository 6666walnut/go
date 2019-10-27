/**
* @Author: Chicken dishes
* @Date: 2019/10/26 15:59
 */

package main

import "fmt"

// 使用值接受者和指针接受者的区别
type animal interface {
	move()
	eat(someting string)
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
		name: "tom",
		feet: 4,
	}
	c2 := &cat{
		name: "tong",
		feet: 4,
	}
	a1 = &c1
	fmt.Printf("类型:%T 值:%v \n",a1,a1)
	a1 = c2
	fmt.Printf("类型:%T 值:%v \n",a1,a1)
}
