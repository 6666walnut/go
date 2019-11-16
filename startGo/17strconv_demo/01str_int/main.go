/**
* @Author: Chicken dishes
* @Date: 2019/10/31 7:59
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := int(2345)
	ret := string(i)
	fmt.Println(ret)
	// 转换为str
	ret1 := fmt.Sprintf("%d\n", i)
	fmt.Printf(ret1)

	str := "1000"
	num, _ := strconv.Atoi(str)
	fmt.Println(num)
	ret2:= strconv.Itoa(int(num))
	fmt.Println(ret2)
}
