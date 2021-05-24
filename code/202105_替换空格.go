/**
* @Author: Chicken dishes
* @Date: 2021/5/23 16:39
**/

package main

import (
	"fmt"
	"strings"
)

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func main() {

	ccc := replaceSpace("We are happy.")
	fmt.Println(ccc)
}
