package controller

import (
	"blog/logger"
	"blog/model"
	"html/template"
	"net/http"
	"strings"
)

func Tag(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/tag/")[1]
	id := model.SelectWhere(model.Category{}, "url=?", "id", url)
	if len((*(id.(*[]model.Category)))) != 0 {
		values := model.SelectPreload(model.Article{}, "Category",
			"*", "category_id=?", (*(id.(*[]model.Category)))[0].ID)
		t, err := template.ParseFiles("view/index/index.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		t.Execute(w, values)
	} else {
		t, err := template.ParseFiles("view/index/404.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		t.Execute(w, nil)
	}

}
