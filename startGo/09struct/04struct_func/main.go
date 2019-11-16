/**
* @Author: Chicken dishes
* @Date: 2019/10/23 8:35
 */
package _4struct_func

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}



func main() {
	p1 := newPerson("n1", 18)
	p2 := newPerson("n2", 19)
	fmt.Println(p1, p2)
}
