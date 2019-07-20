package main

import (
	"MeetingManager/ApiControllers"
	"MeetingManager/RouterControllers"
	_ "MeetingManager/database"
	"MeetingManager/global"
	"github.com/astaxie/beego/session"
	"github.com/gin-gonic/gin"
)

func init() {
	global.GlobalSessions, _ = session.NewManager("memory", &session.ManagerConfig{
		CookieName:      "session",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
	})
	go global.GlobalSessions.GC()

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./static")
	r.GET("/", RouterControllers.LoginController)
	// 个人管理  路由
	personal := r.Group("/personal")
	{
		personal.GET("/notice", RouterControllers.NoticeController)
		personal.GET("/destine", RouterControllers.DestineController)
		personal.GET("/meeting", RouterControllers.MyMeetingController)
		personal.GET("/password", RouterControllers.ModifyPasswordController)
	}
	// 人员管理路由
	personnel := r.Group("/personnel")
	{
		personnel.GET("/department", RouterControllers.DepartmentController)
		personnel.GET("/staffRegister", RouterControllers.StaffRegisterController)
		personnel.GET("/staffApproval", RouterControllers.StaffApprovalController)
		personnel.GET("/staffSearch", RouterControllers.StaffSearchController)
	}
	meeting := r.Group("/meeting")
	{
		meeting.GET("/meetingRoomAdd", RouterControllers.MeetingRoomAddController)
		meeting.GET("/meetingRoomView", RouterControllers.MeetingRoomViewController)
		meeting.GET("/destineMeeting", RouterControllers.DestineMeetingController)
		meeting.GET("/meetingSearch", RouterControllers.MeetingSearchController)
		meeting.GET("/details", RouterControllers.MeetingDetailsController)
		meeting.GET("/myDetails", RouterControllers.MyDetailsController)
		meeting.GET("/roomDetails", RouterControllers.RoomDetailsController)
	}

	api := r.Group("/api")
	{
		api.POST("/login", ApiControllers.LoginController)
		api.GET("/meeting", ApiControllers.MeetingController)
		api.PUT("/meeting", ApiControllers.MeetingStatusChange)
		api.POST("/meeting", ApiControllers.SaveMeeting)
		api.GET("/department", ApiControllers.GetDepartment)
		api.DELETE("/department", ApiControllers.DeleteDepartment)
		api.POST("/department", ApiControllers.SaveDepartment)
		api.POST("/staff", ApiControllers.SaveStaffController)
		api.GET("/staff", ApiControllers.GetStaffRouter)
		api.PUT("/staff", ApiControllers.ChangeStaffStatus)
		api.DELETE("/staff", ApiControllers.DeleteStaff)
		api.POST("/meetingRoom", ApiControllers.SaveMeetingRoom)
		api.GET("/meetingRoom", ApiControllers.GetMeetingRoomRouter)
		api.PUT("/meetingRoom", ApiControllers.UpdateMeetingRoom)
	}
	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}

}
