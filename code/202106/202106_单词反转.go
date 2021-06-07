/**
* @Author: Chicken dishes
* @Date: 2021/6/5 17:38
**/

package main

import (
	"strings"
)

/**
 * 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。
 * 例如输入字符串"I am a student. "，则输出"student. a am I"。
 */

func reverseWords(s string) string {

	arr := strings.Split(s, " ")
	var res []string

	arrLen := len(arr)
	for k := range arr {
		if arr[arrLen-k-1] == "" {
			continue
		}
		res = append(res, arr[arrLen-k-1])
	}
	return strings.Join(res, " ")
}

func main() {
	reverseWords("  hello world!  ")
}
