package ApiControllers

import (
	"MeetingManager/database"
	"MeetingManager/global"
	"MeetingManager/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func GetMeetingRoomRouter(context *gin.Context) {
	tp := context.Query("type")
	switch tp {
	case "0":
		GetMeetingRoom(context)
	case "1":
		GetMeetingRoomByCondition(context)
	}
}

func SaveMeetingRoom(context *gin.Context) {
	value := make(map[string]interface{})
	data, _ := ioutil.ReadAll(context.Request.Body)
	err := json.Unmarshal(data, &value)
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	capacity, err := strconv.Atoi(value["roomCapacity"].(string))
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	err = database.Insert(&model.MeetingRoom{
		RoomID:       value["roomId"].(string),
		RoomName:     value["roomName"].(string),
		RoomCapacity: capacity,
		RoomStatus:   int(value["roomStatus"].(float64)),
		RoomRemarks:  value["roomRemarks"].(string),
	})
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_DATABASE_CANT_INSERT,
			"code":    global.ERROR_DATABASE_CANT_INSERT,
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

func GetMeetingRoom(context *gin.Context) {
	data := database.QueryAll(&model.MeetingRoom{})
	if len(*data.(*[]model.MeetingRoom)) <= 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
			"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
		})
		return
	}
	context.JSON(global.HTTP_STATUS_OK, gin.H{
		"status":  global.HTTP_STATUS_OK,
		"data":    data,
		"message": "",
		"code":    global.NO_ERROR,
	})
}

func GetMeetingRoomByCondition(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	data := database.QueryObject(false, &model.MeetingRoom{}, id)
	if data.(*model.MeetingRoom).ID == 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
			"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
		})
		return
	}
	context.JSON(global.HTTP_STATUS_OK, gin.H{
		"status":  global.HTTP_STATUS_OK,
		"data":    data,
		"message": "",
		"code":    global.NO_ERROR,
	})

}

func UpdateMeetingRoom(context *gin.Context) {
	value := make(map[string]interface{})
	data, _ := ioutil.ReadAll(context.Request.Body)
	err := json.Unmarshal(data, &value)
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	status, err := strconv.Atoi(value["RoomStatus"].(string))
	capacity, err := strconv.Atoi(value["RoomCapacity"].(string))
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
	err = database.Update(&model.MeetingRoom{
		ID:           int(value["ID"].(float64)),
		RoomName:     value["RoomName"].(string),
		RoomRemarks:  value["RoomRemarks"].(string),
		RoomStatus:   status,
		RoomCapacity: capacity,
	})
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
			"message": global.ERROR_STRING_DATABASE_CANT_UPDATE,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		return
	}
}
