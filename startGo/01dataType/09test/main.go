/**
* @Author: Chicken dishes
* @Date: 2019/10/16 9:08
 */

package main

import "fmt"

func main() {
	myStr := "hello黄鹤楼"
	//s2 := []rune(myStr)
	strSlice := []string{}
	for _, num := range myStr {
		if num <= 40869 && num >= 19968 {
			strSlice = append(strSlice, string(num))
		}
	}
	fmt.Println(strSlice)
	fmt.Println(len(strSlice))
}
