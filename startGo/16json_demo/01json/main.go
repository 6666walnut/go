/**
* @Author: Chicken dishes
* @Date: 2019/10/31 6:49
 */

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"tan","age":1}`
	var p person
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age)
}
