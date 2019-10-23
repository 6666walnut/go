/**
* @Author: Chicken dishes
* @Date: 2019/10/23 8:54
 */
package main

import (
	"fmt"
)

/*
学生管理系统
写一个系统能查看\新增学生\删除学生
*/

type students struct {
	data map[int64]*student
}

type student struct {
	num  int64
	name string
}

func (a students) showStudent() {
	fmt.Println("show")
	for k, v := range a.data {
		fmt.Printf("学号:%d 名字:%s\n", k, v.name)
	}
}

func (a students) addStudent() {
	fmt.Println("add")
	var s student
	fmt.Println("请输入学号:")
	fmt.Scanln(&s.num)
	fmt.Println("请输入姓名:")
	fmt.Scanln(&s.name)
	fmt.Println(s)
	a.data[s.num] = &s
}

func (a students) delStudent() {
	fmt.Println("del")
	fmt.Println("请输入学号:")
	var num int64
	fmt.Scanln(&num)
	delete(a.data, num)
}

func main() {
	allstudent := students{
		data: make(map[int64]*student, 48),
	}

	var choice int
	for {
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			allstudent.showStudent()
		case 2:
			allstudent.addStudent()
		case 3:
			allstudent.delStudent()
		default:
			fmt.Println(choice)
			goto xx
		}
	}
xx:
	fmt.Println("退出")

}
