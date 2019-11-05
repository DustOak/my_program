package controller

import (
	"blog/logger"
	"blog/model"
	"html/template"
	"net/http"
)

func B_Comment(w http.ResponseWriter, r *http.Request) {
	if SessionManager.CheckIsExist(SessionManager.CookieIsExist(r, "token")) {
		t, err := template.ParseFiles("view/backup/comments.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		value := model.SelectAll(model.Comment{}, "")
		t.Execute(w, value)
	} else {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}
}
