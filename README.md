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
setting.Logger.Info("string 类型")
setting.Logger.Info("string 类型",interface{}")

说明：setting.Logger 会得到一个grayLog的实例，后期会支持app.ini的参数配置，得到不同的实例,不需要额外的字段，可使用setting.Logger.WithField()生成实例
```
# Code码使用
```
"aos/errors"
eg:
errors.SYSERR // code码
errors.GetInfo()[errors.SYSERR] // code 对应的值
TODO:进度封装，方便使用
```
# Gin Session Middleware
```GO
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    store := sessions.NewCookieStore([]byte("secret")) // Cookie
    // store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret")) // Redis 
	// store := sessions.NewMemcacheStore(memcache.New("localhost:11211"), "", []byte("secret"))
    // Mongo
    // session, err := mgo.Dial("localhost:27017/test")
	// if err != nil {
	// 	// handle err
	// }
	// c := session.DB("").C("sessions")
	// store := sessions.NewMongoStore(c, 3600, true, []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
```

# Sentry
```GO
package main

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

func init() {
	raven.SetDSN("https://<key>:<secret>@app.getsentry.com/<project>")
}

func main() {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	// only send crash reporting
	// r.Use(sentry.Recovery(raven.DefaultClient, true))
	r.Run(":8080")
}
```

# TODO list
- [x] Panic 处理
- [x] sentry 日志
- [x] 加上Swagger
- [x] 支持Cors处理
- [ ] SQL 驱动与ORM选取
- [ ] SQL 日志打印到graylog
- [x] 输出日志打印到graylog
- [ ] Http请求
- [x] Session
- [x] X-Response-ID
- [ ] Consul 读取
- [ ] Redis 封装
- [ ] DDD设计实现
- [x] 状态码统一管理
