/**
* @Author: Chicken dishes
* @Date: 2019/10/31 7:06
 */

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// mysql配置文件
type MysqlConfig struct {
	Address  string
	Port     int
	Username string
	Password string
}

func loadIni(fileName string, data interface{}) (error) {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err := errors.New("data  should be a pointer.")
		return err
	}
	if t.Elem().Kind() != reflect.Struct {
		err := errors.New("*data  should be a struct.")
		return err
	}
	// 1.读文件
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Println(lineSlice)
	// 2.一行一样的读
	for index, line := range lineSlice {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if line[0] != '[' || line[len(line)-1] != ']' {
			err = fmt.Errorf("line:%d syntax error",index)
		}
	}
	// 2.1 如果是注释就跳过
	// 2.2 如果是[]
	// 2.3 如果不是[开头就是=分割的键值对]
	return nil
}

func main() {
	var mv MysqlConfig
	//var x =new(int)
	err := loadIni("./conf.ini", &mv)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mv.Address, mv.Password)
}
