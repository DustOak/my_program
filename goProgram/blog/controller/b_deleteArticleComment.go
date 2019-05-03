package controller

import (
	"blog/logger"
	"blog/model"
	"html"
	"io"
	"net/http"
	"strconv"
)

func B_DeleteArticleComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(html.EscapeString(r.FormValue("id")))
	if err != nil {
		logger.ErrLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = model.Delete(&model.ArticleComment{
		ID: id,
	})
	if err != nil {
		logger.ErrLog.Println(err)
		io.WriteString(w, "-1")
		return
	}

}
