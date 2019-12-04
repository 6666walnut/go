## 如何使用modules

1. golang 版本升级到1.11

2. 设置 GO111MODULE 

     *  GO111MODULE=off  ：  go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。 
     *  GO111MODULE=on ： go命令行会使用modules，而一点也不会去GOPATH目录下查找。 
     *  GO111MODULE=auto ： 默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
          * 当前目录在GOPATH/src之外且该目录包含go.mod文件
          * 当前文件在包含go.mod文件的目录下面

3. go mod 包管理

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
   