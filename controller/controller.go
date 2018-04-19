package controller

import (
	"aos/pkg/errors"
	"aos/pkg/utils"
	"fmt"

	"aos/persistence"
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
func (myc *TestApi) GetS(c *gin.Context) {

	// utils.RedisHandle.SetData("test1", "hhhhh", 0)
	// utils.HttpHandle.Debug = true
	data, err := utils.HttpHandle.Get("http://t-goada.gaodun.com/homework/oneStatisticsInfo?homework_id=33&is_complete_all=1", nil, nil)
	fmt.Println(utils.RedisHandle.GetData("test1"))
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

func (myc *TestApi) GetDbTest(c *gin.Context) {
	var subjectModel = persistence.NewGdSubjectFacFac()
	subjectModel.FindAll("")
	myc.ServerJSON(c, subjectModel.RowsSlicePtr, errors.SUCCESSSTATUS)
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
