package controller

import (
	"blog/logger"
	"blog/model"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

func B_ArticleDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := template.HTMLEscapeString(template.JSEscapeString(r.FormValue("id")))
	if id == "" {
		io.WriteString(w, "-1")
		return
	} else {
		iD, err := strconv.Atoi(id)
		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
		err = model.Delete(&model.Article{
			ID: iD,
		})
		if err != nil {
			logger.ErrLog.Println(err)
			io.WriteString(w, "-1")
			return
		}
		io.WriteString(w, "1")
	}

}
