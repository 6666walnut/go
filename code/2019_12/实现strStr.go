/**
* @Author: Chicken dishes
* @Date: 2019/12/31 21:02
 */
package main

//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	lenN := len(needle)
	lenH := len(haystack)
	for i := 0; lenH-lenN >= i; i++ {
		if haystack[i:i+lenN] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "aaaaa"
	needle := "bba"

	strStr(haystack, needle)
}
