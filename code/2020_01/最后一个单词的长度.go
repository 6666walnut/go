/**
* @Author: Chicken dishes
* @Date: 2020/1/7 21:42
 */
package main

import "fmt"

//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
//
//如果不存在最后一个单词，请返回 0 。

func lengthOfLastWord(s string) int {
	sLen := len(s)
	k1 := 0
	k2 := 0
	for i := sLen - 1; i >= 0; i-- {
		if k1 == 0 && s[i] != 32 {
			k1 = i
		}
		if k1 != 0 && s[i] == 32 {
			k2 = i
			break
		}
	}
	if sLen != 0 && k2 == 0 && s[0] != 32 {
		k2--
	}
	return k1 - k2
}

func main() {
	s := ""
	myLen := lengthOfLastWord(s)
	fmt.Println(myLen)
}
