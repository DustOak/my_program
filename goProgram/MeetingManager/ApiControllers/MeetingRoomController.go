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
