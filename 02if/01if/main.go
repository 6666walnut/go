/**
* @Author: Chicken dishes
* @Date: 2019/10/17 8:16
 */

package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("开业")
	} else {
		fmt.Println("回家")
	}

	if ageN := 19; ageN > 18 {
		fmt.Println("开业")
	} else {
		fmt.Println("回家")
	}

}
