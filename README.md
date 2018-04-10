# godep 使用
```go
安装 go get github.com/tools/godep

godep save 将项目中使用到的第三方库复制到项目的vendor目录下

godep restore godep会按照Godeps/Godeps.json内列表，依次执行go get -d -v 来下载对应依赖包到GOPATH路径下

说明：关于无法安装golang.org下的库时， You can also manually git clone the repository to $GOPATH/src/golang.org/x/sys
```

#grpc 使用
```go
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

// protoc --go_out=plugins=grpc:. helloworld.proto

go get google.golang.org/grpc 翻墙 🤢

在 github 可以找到源码，下载后复制到对应目录即可的： 
google.golang.org/grpc 对应的代码地址在： https://github.com/grpc/grpc-go 
google.golang.org/cloud/compute/metadata 对应的代码地址在： https://github.com/GoogleCloudPlatform/gcloud-golang 
golang.org/x/oauth2 对应的代码地址在： https://github.com/golang/oauth2 
golang.org/x/net/context 对应的代码地址在： https://github.com/golang/net 
这些包的源码也可以通过 http://gopm.io/ 或者 http://golangtc.com/download/package 进行下载.

ERROR:package google.golang.org/genproto/googleapis/rpc/status: unrecognized import path "google.golang.org/genproto/googleapis/rpc/status" (https fetch: Get https://google.golang.org/genproto/googleapis/rpc/status?go-get=1: unexpected EOF)
解决方案：
1、cd $GOPATH/src/google.golang.org
2、git clone https://github.com/google/go-genproto
3、mv -f go-genproto  genproto

```