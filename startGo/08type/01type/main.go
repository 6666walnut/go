/**
* @Author: Chicken dishes
* @Date: 2019/10/22 9:14
 */
package main

import "fmt"

type myInt int

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T \n", n)
}
