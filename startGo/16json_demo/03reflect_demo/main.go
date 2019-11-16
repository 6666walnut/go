/**
* @Author: Chicken dishes
* @Date: 2019/10/31 7:12
 */

package main

import (
	"fmt"
	"reflect"
)

type myInt int64

type person struct {
	name string
	age  int
}

type book struct {
	title string
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v \n", t.Name(), t.Kind())
}

func main() {
	var a *float32
	var b myInt
	var c rune

	reflectType(a)
	reflectType(b)
	reflectType(c)

	var p = person{
		name: "nb",
		age:  18,
	}
	var book1 = book{title: "book"}
	reflectType(p)
	reflectType(book1)
}
