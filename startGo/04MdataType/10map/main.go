/**
* @Author: Chicken dishes
* @Date: 2019/10/19 8:57
 */

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var m1 map[string]int
	fmt.Println(m1)               // 还没有初始化
	m1 = make(map[string]int, 10) //与共鸣

	m1["age"] = 18
	m1["li"] = 19
	fmt.Println(m1)
	fmt.Println(m1["age"])
	v, ok := m1["tt"]
	if !ok {
		fmt.Println(v)
	}
	for k, v := range m1 {   // 遍历
		fmt.Println(k, v)
	}
	delete(m1,"age") // 删除
	fmt.Println(m1)

}

func Sort()  {

		rand.Seed(time.Now().Unix())
		var scoreMap = make(map[string]int, 200)
		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("name%02d", i)
			value := rand.Intn(100)
			scoreMap[key] = value
		}
		fmt.Println(scoreMap)

		var keys = make([]string, 0, 200)
		for key := range scoreMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key,scoreMap[key])
		}

}
