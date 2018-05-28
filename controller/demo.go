package controller

import (
	"aos/pkg/errors"
	"aos/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestApi struct {
	Name string
	Base
}

func NewDemoController() *TestApi {
	var c = &TestApi{Name: "mynameisstone"}
	return c
}

// @Summary 获取S
// @Produce  json
// @Param access_key path strinbee true "秘钥KEY"
// @Success 200 {string} json "{"status": 1,"message": "","result": {"access_key": "xxx","access_secret": ""}}"
// @Router /secret/{access_key} [get]
func (myc *TestApi) GetS(c *gin.Context) {

	// utils.RedisHandle.SetData("test1", "hhhhh", 0)
	// utils.HttpHandle.Debug = true
	data, err := utils.HttpHandle.Get("http://t-goada.gaodun.com/homework/oneStatisticsInfo?homework_id=33&is_complete_all=1", nil, nil)
	fmt.Println(utils.GetRedisHandle().GetData("test1"))
	fmt.Println(err)
	// authentication := CreateSecretFromRequest(c)
	// fmt.Println("Access key is " + authentication.AccessKey)
	// fmt.Println("Access secret is " + authentication.AccessSecret)
	// _, err := secretServiceFacade.Authenticate(authentication)
	// if nil != err {
	// 	fmt.Println(err)
	// }
	myc.ServerJSON(c, data, errors.SUCCESSSTATUS)
	// c.JSON(200, ResponseObject{
	// 	1,
	// 	"",
	// 	authentication,
	// })
}

func (myc *TestApi) setName(name string) {
	// println(reflect.TypeOf(new(project_service_impl.ProjectServiceStruct)))
	// setting.GinLogger(c).Info("abc")
	// myc.Name = "stone"
	myc.Name = name
}

func (myc *TestApi) TestGraylog(c *gin.Context) {
	// println(reflect.TypeOf(new(project_service_impl.ProjectServiceStruct)))
	// setting.GinLogger(c).Info("abc")
	// myc.Name = "stone"
	myc.setName("我啊")
	fmt.Print(myc.Name)
}

func (myc *TestApi) TestName(c *gin.Context) {
	// println(reflect.TypeOf(new(project_service_impl.ProjectServiceStruct)))
	// setting.GinLogger(c).Info("abc")
	// myc.Name = "张三"
	fmt.Print(myc.Name)
}
