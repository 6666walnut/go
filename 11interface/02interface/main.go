/**
* @Author: Chicken dishes
* @Date: 2019/10/26 13:00
 */

package main

import "fmt"

type car interface {
	run()
}

type aodi struct {
	brand string
}

type bm struct {
	brand string
}

func (f aodi) run() {
	fmt.Println(f.brand)
}
func (f bm) run()  {
	fmt.Println(f.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = aodi{
		"A4",
	}
	var f2 = bm{brand: "3"}
	drive(f1)
	drive(f2)
}
