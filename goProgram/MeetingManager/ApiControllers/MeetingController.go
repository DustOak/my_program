package ApiControllers

import (
	"MeetingManager/database"
	"MeetingManager/global"
	"MeetingManager/model"
	"MeetingManager/tools"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

// code 0 查看某个会议 case 1  查看未来七天会议 case 2 查看我的预定会议 3获取今天会议 4将要参加的会议 5搜索会议
func MeetingController(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("<MeetingController>异常 错误信息:[", err.(error).Error(), "]")
		}
	}()
	switch context.Query("type") {
	case "0":
		id, err := strconv.Atoi(context.Query("id"))
		if err != nil {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    gin.H{},
				"message": global.ERROR_STRING_PRAM_ILLEGAL,
				"code":    global.ERROR_PRAM_ILLEGAL})
			return
		}
		GetMeeting(context, id)
	case "1":
		GetSevenDayMeetings(context)
	case "2":
		GetMyBookMeetings(context)
	case "3":
		GetTodayMeetings(context)
	case "4":
		GetIJoinMeetings(context)
	case "5":
		SearchMeeting(context)
	default:
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL})
		return
	}
}

func GetMyBookMeetings(context *gin.Context) {
	sess, _ := global.GlobalSessions.SessionStart(context.Writer, context.Request)
	session := sess.Get("session")
	if session == nil || session.(*model.Personnel).ID == 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_SESSION_NOT_FOUNT,
			"code":    global.ERROR_SESSION_NOT_FOUNT,
		})
	} else {
		a := database.QueryPreloadCondition(&model.Meeting{}, " MEETING_STATUS =1 AND BOOK_PERSONNEL =?", session.(*model.Personnel).ID)
		if len(*a.(*[]model.Meeting)) <= 0 {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    "",
				"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
				"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
			})
		} else {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    a,
				"message": "",
				"code":    global.NO_ERROR,
			})
		}
	}
}

func GetSevenDayMeetings(context *gin.Context) {
	sess, _ := global.GlobalSessions.SessionStart(context.Writer, context.Request)
	session := sess.Get("session")
	if session == nil || session.(*model.Personnel).ID == 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_SESSION_NOT_FOUNT,
			"code":    global.ERROR_SESSION_NOT_FOUNT,
		})
	} else {
		a := database.QueryPreloadCondition(&model.Meeting{}, "(start_time<= timestampadd(day, 7, now())) AND MEETING_STATUS =1    AND book_personnel= ? ", 1)
		if len(*a.(*[]model.Meeting)) <= 0 {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    "",
				"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
				"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
			})
		} else {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    a,
				"message": "",
				"code":    global.NO_ERROR,
			})
		}
	}
}

func GetTodayMeetings(context *gin.Context) {
	sess, _ := global.GlobalSessions.SessionStart(context.Writer, context.Request)
	session := sess.Get("session")
	if session == nil || session.(*model.Personnel).ID == 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_SESSION_NOT_FOUNT,
			"code":    global.ERROR_SESSION_NOT_FOUNT,
		})
	} else {
		a := database.QueryPreloadCondition(&model.Meeting{}, "(to_days(start_time) = to_days(now())) AND book_personnel IN (?)", session.(*model.Personnel).ID)
		if len(*a.(*[]model.Meeting)) <= 0 {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    "",
				"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
				"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
			})
		} else {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    a,
				"message": "",
				"code":    global.NO_ERROR,
			})
		}
	}
}

func GetMeeting(context *gin.Context, id int) {
	sess, _ := global.GlobalSessions.SessionStart(context.Writer, context.Request)
	session := sess.Get("session")
	if session == nil || session.(*model.Personnel).ID == 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_SESSION_NOT_FOUNT,
			"code":    global.ERROR_SESSION_NOT_FOUNT,
		})
		return
	} else {
		meeting := database.QueryObject(true, &model.Meeting{}, id)
		if meeting.(*model.Meeting).ID <= 0 {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    "",
				"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
				"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
			})
		} else {
			personnels := database.Query(&model.Personnel{}, "SELECT P.* FROM PERSONNEL AS P ,"+
				"MEETING_PERSONNEL AS M WHERE M.MEETING_ID =? AND P.ID=M.PERSONNEL_ID", id)
			a := make(map[string]interface{})
			a["meetings"] = meeting
			a["personnels"] = personnels
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    a,
				"message": "",
				"code":    global.NO_ERROR,
			})
		}
	}
}

func GetIJoinMeetings(context *gin.Context) {
	sess, _ := global.GlobalSessions.SessionStart(context.Writer, context.Request)
	session := sess.Get("session")
	if session == nil || session.(*model.Personnel).ID == 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_SESSION_NOT_FOUNT,
			"code":    global.ERROR_SESSION_NOT_FOUNT,
		})
	} else {
		meeting := database.Query(&model.Meeting{}, "SELECT MM.* FROM MEETING_PERSONNEL AS M,MEETING AS MM WHERE  MM.ID=M.MEETING_ID AND MM.MEETING_STATUS =1 AND M.PERSONNEL_ID=?", 1)
		if len(*meeting.(*[]model.Meeting)) <= 0 {
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    "",
				"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
				"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
			})
		} else {
			meetingRoom := database.QueryAll(&model.MeetingRoom{})
			personnel := database.QueryAll(&model.Personnel{})
			for i := 0; i < len(*meeting.(*[]model.Meeting)); i++ {
				for j := 0; j < len(*meetingRoom.(*[]model.MeetingRoom)); j++ {
					if (*meeting.(*[]model.Meeting))[i].RID == (*meetingRoom.(*[]model.MeetingRoom))[j].ID {
						(*meeting.(*[]model.Meeting))[i].MeetingRoom = (*meetingRoom.(*[]model.MeetingRoom))[j]
						break
					}
				}
				for j := 0; j < len(*personnel.(*[]model.Personnel)); j++ {
					if (*meeting.(*[]model.Meeting))[i].BookPersonnel == (*personnel.(*[]model.Personnel))[j].ID {
						(*meeting.(*[]model.Meeting))[i].BP = (*personnel.(*[]model.Personnel))[j]
						break
					}
				}
			}
			context.JSON(global.HTTP_STATUS_OK, gin.H{
				"status":  global.HTTP_STATUS_OK,
				"data":    meeting,
				"message": "",
				"code":    global.NO_ERROR,
			})
		}
	}
}

func MeetingStatusChange(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_PRAM_ILLEGAL,
			"code":    global.ERROR_PRAM_ILLEGAL,
		})
		log.Fatalln(err)
	}
	err = database.Update(&model.Meeting{ID: id, MeetingStatus: -1})
	if err != nil {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_DATABASE_CANT_UPDATE,
			"code":    global.ERROR_DATABASE_CANT_UPDATE,
		})
	}
	context.JSON(global.HTTP_STATUS_OK, gin.H{
		"status":  global.HTTP_STATUS_OK,
		"data":    gin.H{},
		"message": "",
		"code":    global.NO_ERROR,
	})
}

func SearchMeeting(context *gin.Context) {
	var (
		bookDateMax     = context.Query("bookDateMax")
		bookDateMin     = context.Query("bookDateMin")
		meetingBookName = context.Query("meetingBookName")
		meetingDateMax  = context.Query("meetingDateMax")
		meetingDateMin  = context.Query("meetingDateMin")
		meetingName     = context.Query("meetingName")
		meetingRoomID   = context.Query("meetingRoomID")
		bookDateSql     = " and m.BOOK_TIME >= '" + bookDateMin + "' and  m.BOOK_TIME <= '" + bookDateMax + "'"
		meetingDate     = " and m.START_TIME >= '" + meetingDateMin + "' and  m.START_TIME <= '" + meetingDateMax + "'"
		meetingBN       = " and p.PERSONNEL_NAME=  '" + meetingBookName + "'"
		meetingRI       = " and mr.ROOM_ID='" + meetingRoomID + "'"
		meetingN        = " and m.NAME like '%" + meetingName + "%'"
	)

	sql := "select m.* from MEETING as m, PERSONNEL as p ,MEETINGROOM as mr  where  1=1 "
	if strings.TrimSpace(bookDateMax) != "" && strings.TrimSpace(bookDateMin) != "" {
		sql += bookDateSql
	}
	if strings.TrimSpace(meetingDateMax) != "" && strings.TrimSpace(meetingDateMin) != "" {
		sql += meetingDate
	}
	if strings.TrimSpace(meetingBookName) != "" {
		sql += meetingBN
	}
	if strings.TrimSpace(meetingName) != "" {
		sql += meetingN
	}
	if strings.TrimSpace(meetingRoomID) != "" {
		sql += meetingRI
	}
	result := database.Query(&model.Meeting{}, sql)
	if len(*result.(*[]model.Meeting)) <= 0 {
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    gin.H{},
			"message": global.ERROR_STRING_DATEBASE_DATA_NOT_FOUND,
			"code":    global.ERROR_DATEBASE_DATA_NOT_FOUND,
		})
		return
	} else {
		personnel := database.QueryAll(&model.Personnel{})
		meetingRoom := database.QueryAll(&model.MeetingRoom{})
		for i := 0; i < len(*result.(*[]model.Meeting)); i++ {
			for j := 0; j < len(*personnel.(*[]model.Personnel)); j++ {
				if (*result.(*[]model.Meeting))[i].BookPersonnel == (*personnel.(*[]model.Personnel))[j].ID {
					(*result.(*[]model.Meeting))[i].BP = (*personnel.(*[]model.Personnel))[j]
					break
				}
			}
			for j := 0; j < len(*meetingRoom.(*[]model.MeetingRoom)); j++ {
				if (*result.(*[]model.Meeting))[i].RID == (*meetingRoom.(*[]model.MeetingRoom))[j].ID {
					(*result.(*[]model.Meeting))[i].MeetingRoom = (*meetingRoom.(*[]model.MeetingRoom))[j]
					break
				}
			}
		}
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    result,
			"message": "",
			"code":    global.NO_ERROR,
		})
	}

}

func SaveMeeting(context *gin.Context) {
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
		return
	}
	fmt.Println(value)
	st, err := time.Parse("2006-01-02 15:04:05", value["meetingST"].(string))
	et, err := time.Parse("2006-01-02 15:04:05", value["meetingET"].(string))
	err = database.Insert(&model.Meeting{
		RID:                int(value["meetingRoomID"].(float64)),
		Name:               value["meetingName"].(string),
		MeetingStatus:      1,
		BookNumber:         int(value["meetingRoomID"].(float64)),
		StartTime:          tools.LocalTime{Time: st},
		EndTime:            tools.LocalTime{Time: et},
		BookTime:           tools.LocalTime{Time: time.Now()},
		MeetingDescription: value["meetingRemaker"].(string),
	})
	if err != nil {
		log.Println(err)
		context.JSON(global.HTTP_STATUS_OK, gin.H{
			"status":  global.HTTP_STATUS_OK,
			"data":    "",
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
