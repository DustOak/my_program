package controller

import (
	"blog/logger"
	"blog/model"
	"net/http"
	"strings"
	"text/template"
)

func Search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	value := template.HTMLEscapeString(template.JSEscapeString(strings.TrimSpace(r.Form.Get("s"))))
	if value == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	} else {
		values := model.SelectPreload(model.Article{}, "Category", "*",
			"title like ?", "%"+value+"%")
		t, err := template.ParseFiles("view/index/index.html")
		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		t.Execute(w, values)
	}
}
