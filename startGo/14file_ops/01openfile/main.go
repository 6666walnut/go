/**
* @Author: Chicken dishes
* @Date: 2019/10/26 16:54
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 自定义读文件
func openFile() {
	fileObj, err := os.Open("src/goStart/14file_ops/01openfile/main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
	}
	defer fileObj.Close()
	for {
		var tmp [128]byte
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Printf("")
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// 第三方库读文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("src/goStart/14file_ops/01openfile/main.go")
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
	}
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v", err)
			return
		}
		fmt.Println(line)
	}

}

func readFrombyIoutil() {
	ret, err := ioutil.ReadFile("src/goStart/14file_ops/01openfile/main.go")
	if err != nil {
		fmt.Printf("ioutil is failed:%s", err)
	}
	fmt.Println(string(ret))
}

func main() {
	readFrombyIoutil()
}
