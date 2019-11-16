/**
* @Author: Chicken dishes
* @Date: 2019/10/31 6:53
 */

package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int = 6
	reflectType(b)
}
