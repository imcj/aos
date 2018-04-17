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
$ cd $GOPATH/src/aos
swag init 
地址：
http://127.0.0.1:6001/swagger/index.html

```
# Logger 使用
```
暂时支持 graylog
配置conf/app.ini 的log配置 LOG_FIELDS：打印到 graylog 的查询字段
Level = enum [-1,0,1,2,3,4] => ["all","debug","info","warn","error","fatal"]

引入 "aos/pkg/setting"
eg:（Debug、Info、Warn、Error、Fatal）、（Debugf、Infof、Warnf、Errorf、Fatalf）
setting.GrayLog(map[string]interface{}{"what": "I am a tester"}).Info("string 类型")
setting.GrayLog(map[string]interface{}{"what": "I am a tester"}).Info("string 类型",interface{}")

说明：setting.GrayLog(map[string]interface{}{"what": "I am a tester"}) 会得到一个grayLog的实例，后期会支持app.ini的参数配置，得到不同的实例,不需要额外的字段，可使用setting.GrayLog(nil)生成实例
```
# Code码使用
```
"aos/errors"
eg:
errors.SYSERR // code码
errors.GetInfo()[errors.SYSERR] // code 对应的值
TODO:进度封装，方便使用
```
# TODO list
- [ ] SQL 日志打印到graylog
- [ ] 输出日志打印到graylog
- [ ] Http请求
- [ ] Session
- [ ] Request-x-id
- [ ] Consul 读取
- [ ] Redis 封装
- [ ] DDD设计实现