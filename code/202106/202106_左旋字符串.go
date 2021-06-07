/**
* @Author: Chicken dishes
* @Date: 2021/6/5 18:35
**/

package main

import (
	"fmt"
	"strings"
)

/**
 * 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
 * 请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，
 * 该函数将返回左旋转两位得到的结果"cdefgab"。
 */

func reverseLeftWords(s string, n int) string {
	sArr := strings.Split(s, "")
	sArr = append(sArr[n:], sArr[:n]...)
	return strings.Join(sArr, "")
}

func main() {
	s := reverseLeftWords("abcdefg", 2)
	fmt.Println(s)
}
