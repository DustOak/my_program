package RouterControllers

import (
	"MeetingManager/global"
	"github.com/gin-gonic/gin"
)

func LoginController(content *gin.Context) {
	sese, _ := global.GlobalSessions.SessionStart(content.Writer, content.Request)
	a := sese.Get("session")
	if a != nil {
		content.Redirect(global.HTTP_STATUS_REDIRICT, "/personal/notice")
	}
	content.HTML(global.HTTP_STATUS_OK, "login.html", gin.H{})
}
