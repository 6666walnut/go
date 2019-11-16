/**
* @Author: Chicken dishes
* @Date: 2019/10/18 8:13
 */

package main

import "fmt"

func main() {
	s1 := []string{"beijing", "shanghai", "shenzhen"}
	fmt.Printf("s1=%v len:%d cap:%d \n",s1,len(s1),cap(s1))
	s1 = append(s1, "wuhan") // append追加元素,容量不够，跟换内存位置
	fmt.Printf("s1=%v len:%d cap:%d \n",s1,len(s1),cap(s1))

	s1 = append(s1,"hangzhou")
	fmt.Printf("s1=%v len:%d cap:%d \n",s1,len(s1),cap(s1))
}
