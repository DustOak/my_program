package controller

import (
	"blog/logger"
	"blog/model"
	"blog/randomImage"
	"html/template"
	"net/http"
)

func Comment(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index/gustbook.html")
	if err != nil {
		logger.ErrLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value := model.SelectAll(model.Comment{}, "")
	values := make(map[string]interface{})
	values["comment"] = value
	values["headimage"] = randomImage.GetImage()
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t.Execute(w, values)
}
