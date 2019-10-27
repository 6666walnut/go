/**
* @Author: Chicken dishes
* @Date: 2019/10/26 16:27
 */

package main

import "fmt"

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Printf("是个字符串:%s", str)
	}
}
func assign2(a interface{}) {
	fmt.Printf("%T",a)
	switch t := a.(type) {
	case string:
		fmt.Printf("是个字符串，值为:%s\n", t)
	case int:
		fmt.Printf("是个int，值为:%s\n",t)
	case bool:
		fmt.Printf("是个int，值为:%s\n",t)
	}
}

func main() {
	assign(100)
	assign2("100")
	assign2(100)
	assign2(nil)
}
