/**
* @Author: Chicken dishes
* @Date: 2019/10/18 9:07
 */

package main

import "fmt"

func main() {

	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)
}
