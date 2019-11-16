/**
* @Author: Chicken dishes
* @Date: 2019/10/16 8:42
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "'E:\\share\\golang\\src\\goStart'"
	fmt.Println(path)

	s := `
		人年少时不明白忧愁的滋味，
喜欢登高远望	
         `
	fmt.Println(s)

	firstName := "tan"

	lastName := "hh"

	// 字符串合并
	ss := firstName + lastName
	ss1 := fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(ss, ss1)

	// 字符串长度
	strLen := len(ss)
	fmt.Println(strLen)

	// 切割字符串
	pathList := strings.Split(path, "\\")
	fmt.Println(pathList)

	// 包含字符串
	str1 := "123"
	str2 := "12345678123"
	// 参数1包含参数2
	fmt.Println(strings.Contains(str1, str2))
	fmt.Println(strings.Contains(str2, str1))

	// 前/后缀判断
	fmt.Println(strings.HasPrefix(str2,str1))
	fmt.Println(strings.HasSuffix(str2,str1))

	// 子串出现的位子
	fmt.Println(strings.Index(str2,str1))
	fmt.Println(strings.LastIndex(str2,str1))

	// join 操作

	fmt.Println(strings.Join(pathList,"--"))
}
