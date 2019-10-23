/**
* @Author: Chicken dishes
* @Date: 2019/10/23 8:42
 */
package main

import "fmt"

// 需要导入dog包的时候首字母需要大写
type dog struct {
	name string
	age  int
}

func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

// 方法是作用域特点类型的函数
// 接受着receiver类似python的self
func (d dog) cry() {
	fmt.Printf("%s:wangwangwang \n", d.name)
}

func (d dog) growUp() {
	d.age++
}

func (d *dog) growUpN() {
	d.age++
}

func main() {
	d1 := newDog("nb",12)
	d1.cry()
	d1.growUp()
	fmt.Println(d1)
	d1.growUpN()
	fmt.Println(d1)
}
