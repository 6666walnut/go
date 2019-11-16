/**
* @Author: Chicken dishes
* @Date: 2019/10/21 8:24
 */

package main

import "fmt"

func main() {
	// Print
	fmt.Print("good")
	fmt.Print("good2")

	// Println 换行
	fmt.Println("good")
	fmt.Println("good2")

	// %T:类型
	// %d:十进制
	// %b:二进制
	// %o:八进制
	// %x:十六进制
	// %c:字符
	// %s:字符串
	// %p:指针
	// %v:值
	// %f:浮点数
	// %f:布尔值

	var m1 = make(map[string]int, 1)
	m1["age"] = 18
	fmt.Printf("%v \n", m1)
	fmt.Printf("%#v \n", m1)

	fmt.Printf("%5.5s\n", "i have a dearm")

	//var s string
	//fmt.Scan(&s)
	//fmt.Println("用户的输入为：", s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scan(&name, &age, &class)
	fmt.Println(name,age,class)
}
