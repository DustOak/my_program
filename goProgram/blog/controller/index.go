package controller

import (
	"blog/logger"
	"blog/model"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Value := model.SelectAll(model.Article{}, "Category")
	t, err := template.ParseFiles("view/index/index.html")
	if err != nil {
		logger.ErrLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t.Execute(w, Value)
}
