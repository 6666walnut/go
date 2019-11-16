/**
* @Author: Chicken dishes
* @Date: 2019/10/16 9:02
 */

package main

import "fmt"

func main() {

	s1 := "黄鹤楼"
	s2 := []rune(s1)
	s2[0] = '红'
	fmt.Println(string(s2))

	c1 := "红"
	c2 := '红' //rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T c4:%T", c3, c4)


	//类型转换
	n := 10
	var f float64
	f = float64(n)
	fmt.Println(f)
	fmt.Printf("%T",f)
}
