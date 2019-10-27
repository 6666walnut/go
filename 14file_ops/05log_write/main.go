/**
* @Author: Chicken dishes
* @Date: 2019/10/27 12:06
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "this is logs")
	fileObj,_ := os.OpenFile("./xxx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprint(fileObj,"this is filebobj")
}
