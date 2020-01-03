/**
* @Author: Chicken dishes
* @Date: 2020/1/3 21:06
 */
package main

import (
	"fmt"
	"strconv"
)

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
