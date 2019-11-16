/**
* @Author: Chicken dishes
* @Date: 2019/10/31 8:46
 */

package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要初始化才能使用

func noBufChannel() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x:=<-ch
		fmt.Println("this is not init num",x)
	}()
	ch <- 10
	wg.Wait()
}

func bufChannel() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x:=<-ch
		fmt.Println("this is  init num",x)
	}()
	//ch <- 10
	wg.Wait()
}
func main() {
	fmt.Println(b)
	b = make(chan int)     // 不带缓冲区通道初始化
	b = make(chan int, 16) // 带缓存区通道初始化
	fmt.Println(b)
	ch := make(chan int, 10)
	ch <- 10
	fmt.Println(ch)
	x := <-ch
	fmt.Println(x)
	noBufChannel()
	bufChannel()
}
