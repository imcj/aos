package controller

import (
	"aos/secret"

	"github.com/gin-gonic/gin"
)

func AddNewSecret(c *gin.Context) {

}

// @Summary 获取S
// @Produce  json
// @Param access_key path string true "秘钥KEY"
// @Success 200 {string} json "{"status": 1,"message": "","result": {"access_key": "xxx","access_secret": ""}}"
// @Router /secret/{access_key} [get]
func GetS(c *gin.Context) {
	// var s map[string]string
	// s["a"] = "2"
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// TODO: 对象依赖配置放到专门的模块
	// var (
	// 	secretDAO           = persistence.NewSecretDAO(client)
	// 	secretServiceFacade = secret.NewSecretServiceFacadeImpl(
	// 		secretDAO,
	// 		secret.NewSecretFactory(),
	// 	)
	// )
	// setting.Logger.Info("xxxxxxxx")
	// setting.GrayLog(map[string]interface{}{"what": "I am a tester"}).Info("22222")
	// setting.GrayLog().Info("33335555")

	authentication := CreateSecretFromRequest(c)
	// fmt.Println("Access key is " + authentication.AccessKey)
	// fmt.Println("Access secret is " + authentication.AccessSecret)
	// _, err := secretServiceFacade.Authenticate(authentication)
	// if nil != err {
	// 	fmt.Println(err)
	// }
	c.JSON(200, ResponseObject{
		1,
		"",
		authentication,
	})
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
