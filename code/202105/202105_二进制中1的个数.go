/**
* @Author: Chicken dishes
* @Date: 2021/5/25 19:51
**/

package main

/**
 * 请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。
 * 例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
 */

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	var res = 0
	for num > 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

func main() {
	hammingWeight(9)
}
