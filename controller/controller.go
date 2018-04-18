package controller

import (
	"aos/errors"
	"aos/utils"
	"fmt"

	"aos/secret"

	"github.com/gin-gonic/gin"
)

type TestApi struct {
	Base
}

func AddNewSecret(c *gin.Context) {

}

// @Summary 获取S
// @Produce  json
// @Param access_key path string true "秘钥KEY"
// @Success 200 {string} json "{"status": 1,"message": "","result": {"access_key": "xxx","access_secret": ""}}"
// @Router /secret/{access_key} [get]
func (tapi *TestApi) GetS(c *gin.Context) {

	// utils.RedisHandle.SetData("test1", "hhhhh", 0)

	fmt.Println(utils.RedisHandle.GetData("test1"))

	authentication := CreateSecretFromRequest(c)
	// fmt.Println("Access key is " + authentication.AccessKey)
	// fmt.Println("Access secret is " + authentication.AccessSecret)
	// _, err := secretServiceFacade.Authenticate(authentication)
	// if nil != err {
	// 	fmt.Println(err)
	// }
	tapi.ServerJSON(c, authentication, errors.SUCCESSSTATUS)
	// c.JSON(200, ResponseObject{
	// 	1,
	// 	"",
	// 	authentication,
	// })
}

func CreateSecretFromRequest(c *gin.Context) secret.Secret {
	accessKey := c.PostForm("access_key")
	if accessKey == "" {
		accessKey = c.Param("access_key")
	}
	accessSecret := c.DefaultQuery("access_secret", "")

	return secret.Secret{
		accessKey,
		accessSecret,
	}
}
