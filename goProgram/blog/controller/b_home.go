package controller

import (
	"blog/logger"
	"blog/model"
	"html/template"
	"net/http"
	"time"
)

func B_Home(w http.ResponseWriter, r *http.Request) {
	token := SessionManager.CookieIsExist(r, "token")
	if SessionManager.CheckIsExist(token) {
		r.ParseForm()
		email := SessionManager.CookieIsExist(r, "email")
		if email == "" {
			http.Redirect(w, r, "/admin", http.StatusFound)
			return
		}
		value := model.SelectWhere(model.Admin{}, "email=?", "id,email,password,lastTime,lastIpAddress", email)
		t, err := template.ParseFiles("view/backup/index.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		temp := (*(value.(*[]model.Admin)))[0]
		temp.LastTime = time.Now().Format("2006-01-02 15:04:05")
		temp.LastIpAddress = r.RemoteAddr
		err = model.Update(&temp)
		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
		t.Execute(w, (*(value.(*[]model.Admin)))[0])
	} else {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}
}
