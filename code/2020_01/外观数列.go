/**
* @Author: Chicken dishes
* @Date: 2020/1/3 21:06
 */
package main

import (
	"fmt"
	"strconv"
)

//「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	num := countAndSay(n-1) + "*"
	count := 1
	myStr := ""
	for i := 0; i < len(num)-1; i++ {
		if num[i] == num[i+1] {
			count++
		} else {
			myStr = myStr + strconv.Itoa(count) + string(num[i])
			count = 1
		}
	}
	return myStr
}

func main() {
	myStr := countAndSay(6)
	fmt.Println(myStr)
}
