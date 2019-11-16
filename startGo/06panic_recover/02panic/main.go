/**
* @Author: Chicken dishes
* @Date: 2019/10/22 8:49
 */
package main

import "fmt"

func f1()  {
	defer func() {
		err := recover()
		fmt.Println("NB:",err)
	}()
	panic("this is err1")
	fmt.Println("f1")
}

func main()  {
	f1()
}
