/**
* @Author: Chicken dishes
* @Date: 2019/10/16 7:56
 */

package main

import "fmt"

// 单一声明
//var name string
//var age int
//var isOk bool

// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	fmt.Print(age)
	fmt.Printf("name:%s", name)
	fmt.Println(isOk)

	// 声明
	var fistName string = "what"
	// 类型推导声明
	var goodName = "where"
	//简短变量声明（函数中可用）
	lastName := "NB"
	fmt.Println(fistName, goodName, lastName)
}
