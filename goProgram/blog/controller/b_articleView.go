package controller

import (
	"blog/logger"
	"blog/model"
	"html"
	"html/template"
	"net/http"
	"strconv"
)

func B_ArticleView(w http.ResponseWriter, r *http.Request) {
	if SessionManager.CheckIsExist(SessionManager.CookieIsExist(r, "token")) {
		r.ParseForm()
		id, err := strconv.Atoi(html.EscapeString(r.FormValue("id")))
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		article := model.Select(&model.Article{}, id, "Category")
		article.(*model.Article).Content = html.UnescapeString((article).(*model.Article).Content)
		articleComment := model.SelectWhere(model.ArticleComment{}, "article_id=?",
			"id,name,content,email,url,date,ipAddress", id)
		cate := model.SelectAll(model.Category{}, "")
		values := make(map[string]interface{})
		values["article"] = article
		values["cate"] = cate
		values["articleComment"] = articleComment
		t, err := template.ParseFiles("view/backup/article.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, values)
	} else {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}
}
