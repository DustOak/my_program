package controller

import (
	"blog/logger"
	"blog/model"
	"html/template"
	"net/http"
)

func B_Article(w http.ResponseWriter, r *http.Request) {
	value := model.SelectAll(model.Article{}, "Category")
	cate := model.SelectAll(model.Category{}, "")
	if SessionManager.CheckIsExist(SessionManager.CookieIsExist(r, "token")) {
		t, err := template.ParseFiles("view/backup/articles.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		values := make(map[string]interface{})
		values["articles"] = value
		values["cate"] = cate
		t.Execute(w, values)
	} else {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}
}
