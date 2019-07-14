package ApiControllers

import (
	"MeetingManager/database"
	"MeetingManager/global"
	"MeetingManager/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetStaffRouter(context *gin.Context) {
	tp := context.Query("type")
	switch tp {
	case "0":
		GetStaff(context)
	case "1":
		SearchStaff(context)
	default:
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL})
	}
}

func SaveStaffController(context *gin.Context) {
	value := make(map[string]interface{})
	data, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		log.Println(err)
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	err = json.Unmarshal(data, &value)
	if err != nil {
		log.Println(err)
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
	}
	if err != nil {
		log.Println(err)
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	err = database.Insert(&model.Personnel{
		PersonnelName:     value["staffName"].(string),
		PersonnelAccount:  value["staffAccount"].(string),
		PersonnelPassword: value["staffPassword"].(string),
		PersonnelEmail:    value["staffEmail"].(string),
		PersonnelPhone:    value["staffPhone"].(string),
		PersonnelStatus:   -1,
		DepartmentID:      int(value["staffDepartment"].(float64)),
	})
	if err != nil {
		log.Println(err)
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_DATABASE_CANT_INSERT,
			"code":    global.ERROR_DATABASE_CANT_INSERT,
		})
	} else {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": "",
			"code":    global.NO_ERROR,
		})
	}
}

func GetStaff(context *gin.Context) {
	data := database.QueryAll(&model.Personnel{})
	if len(*data.(*[]model.Personnel)) <= 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
			"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
		})
	} else {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    data,
			"message": "",
			"code":    global.NO_ERROR,
		})
	}
}

func ChangeStaffStatus(context *gin.Context) {
	id := context.Query("id")
	st := context.Query("status")
	if strings.TrimSpace(id) == "" || strings.TrimSpace(st) == "" {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	ids, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	var status int
	if st != "" {
		status, err = strconv.Atoi(st)
		if err != nil || status > 1 || status < (-2) {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    "",
				"message": global.ERROR_STRING_PRAM_ILLEGAL,
				"code":    global.ERROR_PRAM_ILLEGAL,
			})
			return
		}
	}
	err = database.Update(&model.Personnel{ID: ids, PersonnelStatus: status})
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_DATABASE_CANT_UPDATE,
			"code":    global.ERROR_DATABASE_CANT_UPDATE,
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

func DeleteStaff(context *gin.Context) {
	id := context.Query("id")
	if strings.TrimSpace(id) == "" {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	ids, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	err = database.Delete(&model.Personnel{ID: ids})
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

func SearchStaff(context *gin.Context) {
	name, account, st := context.Query("name"), context.Query("account"), context.Query("status")
	if strings.TrimSpace(name) == "" && strings.TrimSpace(account) == "" && strings.TrimSpace(st) == "" {
		GetStaff(context)
	}
	status, err := strconv.Atoi(st)
	if err != nil {
		log.Println(err)
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	data := database.QueryAll(&model.Personnel{PersonnelName: name, PersonnelAccount: account, PersonnelStatus: status})
	if len(*data.(*[]model.Personnel)) <= 0 {
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
			"data":    data,
			"message": "",
			"code":    global.NO_ERROR,
		})
	}
}
