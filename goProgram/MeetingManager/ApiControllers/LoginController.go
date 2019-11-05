package ApiControllers

import (
	"MeetingManager/database"
	"MeetingManager/global"
	"MeetingManager/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

func LoginController(context *gin.Context) {
	temp := make(map[string]interface{})
	data, _ := ioutil.ReadAll(context.Request.Body)
	err := json.Unmarshal(data, &temp)
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("<LoginController>异常 错误信息:[", err.(error).Error(), "]")
		}
	}()
	if err != nil {
		panic(err)
	}
	account, password := temp["account"], temp["password"]
	if strings.Replace(account.(string), " ", "", -1) == "" || strings.Replace(password.(string), " ", "", -1) == "" {
		context.JSON(200, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.LOGIN_STRING_ACCOUNT_OR_PPASSWORD_ILLEGAL,
			"code":    global.LOGIN_ACCOUNT_OR_PPASSWORD_ILLEGAL,
		})
		return
	}
	sess, _ := global.GlobalSessions.SessionStart(context.Writer, context.Request)
	oj := database.QueryByCondition(&model.Personnel{PersonnelAccount: account.(string), PersonnelPassword: password.(string)})
	if oj.(*model.Personnel).ID != 0 {
		defer sess.SessionRelease(context.Writer)
		err := sess.Set("session", oj)
		if err != nil {
			panic(err)
		}
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.LOGIN_STRING_SUCCESS,
			"code":    global.LOGIN_SUCCESS,
		})
	} else {
		context.JSON(200, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.LOGIN_STRING_ACCOUNT_OR_PASSWORD_WRONG,
			"code":    global.LOGIN_ACCOUNT_OR_PASSWORD_WRONG,
		})
	}
}
