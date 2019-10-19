/**
* @Author: Chicken dishes
* @Date: 2019/10/19 14:37
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	myStr := "how do you do"
	mySlice := strings.Split(myStr, " ")
	myMap := make(map[string]int, 10)
	for _, v := range mySlice {
		myMap[v] = myMap[v] + 1
	}
	fmt.Println(myMap)

}
