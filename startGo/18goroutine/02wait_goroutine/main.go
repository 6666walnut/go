/**
* @Author: Chicken dishes
* @Date: 2019/10/31 8:23
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 5; i++ {
		num1 := rand.Int()
		num2 := rand.Intn(10)
		fmt.Println(num1, num2)
	}

}

var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	fmt.Println(i)
}

func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}

