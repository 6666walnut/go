/**
* @Author: Chicken dishes
* @Date: 2021/5/31 20:36
**/

package main

/**
 * 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
 */

func firstUniqChar(s string) byte {

	var myMap = make(map[int32]int, 0)
	for _, v := range s {
		myMap[v]++
	}
	for _, v := range s {
		if myMap[v] == 1 {
			return byte(v)
		}
	}
	return ' '
}

func main() {
	firstUniqChar("sajd")
}
