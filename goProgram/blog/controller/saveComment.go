package controller

import (
	"blog/logger"
	"blog/model"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func SaveComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var (
		articleId = -1
		err       error
	)
	if r.Form.Get("articleID") != "" {
		articleId, err = strconv.Atoi(r.Form.Get("articleID"))
		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
	}
	author := template.JSEscapeString(template.HTMLEscapeString(r.Form.Get("author")))
	comment := template.JSEscapeString(template.HTMLEscapeString(r.Form.Get("comment")))
	email := template.JSEscapeString(template.HTMLEscapeString(r.Form.Get("email")))
	url := template.JSEscapeString(template.HTMLEscapeString(r.Form.Get("url")))
	headimagepath := template.JSEscapeString(template.HTMLEscapeString(r.Form.Get("HeadImagePath")))
	if articleId != -1 {
		err = model.Save(&model.ArticleComment{
			Name:          author,
			Content:       comment,
			Email:         email,
			Url:           url,
			HeadImagePath: headimagepath,
			ArticleId:     articleId,
			Date:          time.Now().Format("2006-01-02"),
			IpAddress:     r.RemoteAddr,
		})
		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
		http.Redirect(w, r, "/article?id="+strconv.Itoa(articleId), http.StatusFound)
	} else {
		err = model.Save(&model.Comment{
			Name:          author,
			Content:       comment,
			Email:         email,
			Url:           url,
			HeadImagePath: headimagepath,
			Date:          time.Now().Format("2006-01-02"),
			IpAddress:     r.RemoteAddr,
		})
		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
		http.Redirect(w, r, "/comment", http.StatusFound)
	}
}
