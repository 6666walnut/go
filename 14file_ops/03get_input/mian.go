/**
* @Author: Chicken dishes
* @Date: 2019/10/27 11:41
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func userBufio(){
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	s,_ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是:%s\n",s)
}

func main()  {
	userBufio()
}
