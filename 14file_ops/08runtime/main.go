/**
* @Author: Chicken dishes
* @Date: 2019/10/29 9:06
 */

package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("不OK\n")
		return
	}
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {
	f1()
}
