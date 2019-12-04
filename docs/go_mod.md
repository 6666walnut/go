
## 1. 版本要求

* [x] 需要golang 版本 >= 1.11
----
## 2. 设置 GO111MODULE 

1. GO111MODULE=off ：go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。 
2. GO111MODULE=on ：go命令行会使用modules，而一点也不会去GOPATH目录下查找。 
3. GO111MODULE=auto ：默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
     * 当前目录在GOPATH/src之外且该目录包含go.mod文件
     * 当前文件在包含go.mod文件的目录下面
---
## 3. go mod 包管理

| 命令 | 说明 |
| --- | --- |
| download |download modules to local cache(下载依赖包)|
| edit | edit go.mod from tools or scripts(编辑go.mod)|
| graph | print module requirement graph (打印模块依赖图)|
| init | initialize new module in current directory(在当前目录初始化mod)|
| tidy | add missing and remove unused modules(拉取缺少的模块，移除不用的模块)|
| vendor | make vendored copy of dependencies(将依赖复制到vendor下)|
| verify | verify dependencies have expected content (验证依赖是否正确）|
| why | explain why packages or modules are needed(解释为什么需要依赖)|
-----
## 4. 初始化一个新项目

```bash
mkdir new_project
cd new_project
go mod init new_project
```
* go.mod文件一旦创建后，它的内容将会被go toolchain全面掌控。go toolchain会在各类命令执行时，比如go get、go build、go mod等修改和维护go.mod文件。
-----
## 5.go.mod命令

* module  语句指定包的名字（路径）
* require 语句指定的依赖项模块
* replace 语句可以替换依赖项模块
* exclude 语句可以忽略依赖项模块

## 6.依赖添加

1. 创建一个server.go 文件

```golang
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
2. 执行 go run server.go 运行代码会发现 go mod 会自动查找依赖自动下载
3. go.md文件如下：
```
module new_project

go 1.12

require github.com/gin-gonic/gin v1.5.0 // indirect

```
4.其他命令
* go list -m -u all ： 加粗可以升级的包
* go get -u need-upgrade-package ： 升级依赖
* go get -u : 升级所有的依赖