package controller

import (
	"blog/logger"
	"blog/model"
	"html"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"
)

func B_SaveArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := template.HTMLEscapeString(template.JSEscapeString(r.FormValue("id")))
	content := html.EscapeString(r.FormValue("content"))
	title := template.HTMLEscapeString(template.JSEscapeString(r.FormValue("title")))
	tag := template.HTMLEscapeString(template.JSEscapeString(r.FormValue("tag")))
	cid, err := strconv.Atoi(tag)
	if err != nil {
		logger.ErrLog.Println(err)
		io.WriteString(w, "-1")
		return
	}
	if id == "" {
		err = model.Save(&model.Article{
			Title:      title,
			Content:    content,
			CategoryId: cid,
			Date:       time.Now().Format("2006-01-02"),
		})
		if err != nil {
			logger.ErrLog.Println(err)
			io.WriteString(w, "-1")
			return
		}
	} else {
		aid, err := strconv.Atoi(id)
		if err != nil {
			logger.ErrLog.Println(err)
			io.WriteString(w, "-1")
			return
		}
		err = model.Update(&model.Article{
			ID:         aid,
			Title:      title,
			CategoryId: cid,
			Content:    content,
			Date:       time.Now().Format("2006-01-02"),
		})
		if err != nil {
			logger.ErrLog.Println(err)
			io.WriteString(w, "-1")
			return
		}
	}
}
