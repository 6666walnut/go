/**
* @Author: Chicken dishes
* @Date: 2019/10/26 12:47
 */

package main

import "fmt"

type speaker interface {
	speak()
}

type cat struct {
}

type dog struct {

}

func (c cat) speak()  {
	fmt.Println("喵喵喵")
}

func (d dog) speak()  {
	fmt.Println()
}

func da(x speaker)  {
	x.speak()
}

func main()  {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)
}