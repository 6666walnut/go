/**
* @Author: Chicken dishes
* @Date: 2019/10/26 17:15
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file is faile:%s", err)
	}
	defer fileObj.Close()
	// write
	f, err := fileObj.Write([]byte("i am body\n"))
	fmt.Println(f)
	f2, err := fileObj.WriteString("i am body2\n")
	fmt.Println(f2)
}
