# godep 使用
```
安装 go get github.com/tools/godep

godep save 将项目中使用到的第三方库复制到项目的vendor目录下

godep restore godep会按照Godeps/Godeps.json内列表，依次执行go get -d -v 来下载对应依赖包到GOPATH路径下

说明：关于无法安装golang.org下的库时， You can also manually git clone the repository to $GOPATH/src/golang.org/x/sys
```

# 加上Swagger
```
安装 swag
$ go get -u github.com/swaggo/swag/cmd/swag
$ swag -v
依赖golang.org的包
如若无科学上网，可用以下

$ gopm get -g -v github.com/swaggo/swag/cmd/swag
$ cd $GOPATH/src/github.com/swaggo/swag/cmd/swag
$ go install

gopm 安装：$ go get -u github.com/gpmgo/gopm
```
```
安装 gin-swagger
$ go get -u github.com/swaggo/gin-swagger

$ go get -u github.com/swaggo/gin-swagger/swaggerFiles

使用
cd 项目目录
swag init 
```