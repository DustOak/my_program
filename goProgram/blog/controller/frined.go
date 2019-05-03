package controller

import (
	"blog/logger"
	"html/template"
	"net/http"
)

func Friend(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index/link.html")
	if err != nil {
		logger.ErrLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t.Execute(w, nil)
}
