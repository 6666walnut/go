/**
* @Author: Chicken dishes
* @Date: 2019/10/26 16:37
 */

package main

// 禁止循环导入
import (
	"fmt"
	"goStart/13import_package/calc"
)

var x  = 100

const p1  = 3.14

func init()  {
	fmt.Println(x)
}

func main() {
	num := calc.Add(4, 5)
	fmt.Printf("值:%d\n",num)
	fmt.Println(p1)
}
