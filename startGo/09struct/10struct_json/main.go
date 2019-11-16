/**
* @Author: Chicken dishes
* @Date: 2019/10/24 8:56
 */
package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json
// 1.把go语言中的结构体变量 ---> json 格式的字符串
// 2.json格式的字符串 --> go 语言中能识别的结构体变量

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "p1",
		Age:  18,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", string(b))

	// 反序列化
	var p2 person
	str := `{"name":"p2","age":18}`
	_ = json.Unmarshal([]byte(str), &p2)
	fmt.Println(p2)
}
