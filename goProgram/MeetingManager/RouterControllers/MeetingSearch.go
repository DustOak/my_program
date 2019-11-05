package RouterControllers

import (
	"MeetingManager/global"
	"MeetingManager/model"
	"github.com/gin-gonic/gin"
)

func MeetingSearchController(content *gin.Context) {
	sess, _ := global.GlobalSessions.SessionStart(content.Writer, content.Request)
	session := sess.Get("session")
	if session == nil || session.(*model.Personnel).ID == 0 {
		content.Redirect(302, "/")
		return
	}
	content.HTML(global.HTTP_STATUS_OK, "searchmeetings.html", gin.H{})
}
