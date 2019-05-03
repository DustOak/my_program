package controller

import (
	"blog/emailValidate"
	"io"
	"net/http"
	"regexp"
)

func B_SendEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	token := r.FormValue("token")
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])", email); m {
		emailValidate.SendEmail(email, token)
		io.WriteString(w, "发送成功")
		return
	} else {
		io.WriteString(w, "邮箱格式错误")
		return
	}

}
