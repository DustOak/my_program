package ApiControllers

import (
	"MeetingManager/database"
	"MeetingManager/global"
	"MeetingManager/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html"
	"io/ioutil"
	"strconv"
)

func GetDepartment(context *gin.Context) {
	department := database.QueryAll(&model.Department{})
	if len(*department.(*[]model.Department)) <= 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
			"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
		})
		return
	} else {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    department,
			"message": "",
			"code":    global.NO_ERROR,
		})
	}
}

func DeleteDepartment(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	err = database.Exec("UPDATE PERSONNEL SET DEPARTMENT_ID=1 WHERE DEPARTMENT_ID=?", id)
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_DATABASE_DELETE_FAILDE,
			"code":    global.ERROR_DATABASE_DELETE_FAILDE,
		})
		return
	}
	err = database.Delete(&model.Department{ID: id})
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_DATABASE_DELETE_FAILDE,
			"code":    global.ERROR_DATABASE_DELETE_FAILDE,
		})
		return
	}
	context.JSON(global.HTTP_STATUS_OK, gin.H{
		"status":  global.HTTP_STATUS_OK,
		"data":    "",
		"message": "",
		"code":    global.NO_ERROR,
	})
}

func SaveDepartment(context *gin.Context) {
	value := make(map[string]interface{})
	data, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
	}
	err = json.Unmarshal(data, &value)
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
	}
	err = database.Insert(&model.Department{DepartmentName: html.EscapeString(value["departmentname"].(string))})
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_DATABASE_CANT_INSERT,
			"code":    global.ERROR_DATABASE_CANT_INSERT,
		})
		return
	} else {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": "",
			"code":    global.NO_ERROR,
		})
		return
	}

}
