package controller

import (
	"blog/emailValidate"
	"blog/logger"
	"blog/session"
	"html/template"
	"net/http"
	"time"
)

var SessionManager *session.SessionManager

func init() {
	SessionManager = session.NewSessionManager(3600)
	go SessionManager.TimingCleanSession(3000 * time.Second)
}

func B_Index(w http.ResponseWriter, r *http.Request) {
	token := SessionManager.CookieIsExist(r, "token")
	if SessionManager.CheckIsExist(token) {
		http.Redirect(w, r, "/admin/home", http.StatusFound)
		return
	} else {
		t, err := template.ParseFiles("view/backup/login.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		token := emailValidate.GetToken()
		valus := make(map[string]interface{})
		valus["token"] = token
		valus["display"] = "none"
		t.Execute(w, valus)
	}
}
