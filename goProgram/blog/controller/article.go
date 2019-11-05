package controller

import (
	"blog/logger"
	"blog/model"
	"blog/randomImage"
	"html"
	"html/template"
	"net/http"
	"strconv"
)

func Atricle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := template.JSEscapeString(template.HTMLEscapeString(r.FormValue("id")))
	iid, err := strconv.Atoi(id)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	value := model.Select(model.Article{}, iid, "Category")
	a := model.Article{}
	if *(value).(*model.Article) == a {
		t, err := template.ParseFiles("view/index/404.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		t.Execute(w, nil)
		return

	}
	comment := model.SelectWhere(model.ArticleComment{}, "article_id=?",
		"name,content,date,headImagePath", iid)

	t, err := template.ParseFiles("view/index/detail.html")
	if err != nil {
		logger.ErrLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	values := make(map[string]interface{})
	(*(value).(*model.Article)).Content = html.UnescapeString((*(value).(*model.Article)).Content)
	values["article"] = value
	values["headimage"] = randomImage.GetImage()
	values["articleComment"] = comment
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t.Execute(w, values)
}
