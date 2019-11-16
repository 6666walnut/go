/**
* @Author: Chicken dishes
* @Date: 2019/10/31 8:16
 */

package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello world", i)
}

func main() {
	for i := 0; i < 100; i++ {
		//go hello(i)
		go func(i int) {
			fmt.Println("hello world",i)
		}(i)
	}

	fmt.Println("main")
	time.Sleep(time.Second * 1)
}
