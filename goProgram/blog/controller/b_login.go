package controller

import (
	"blog/emailValidate"
	"blog/logger"
	"blog/model"
	"crypto/md5"
	"fmt"
	"html/template"
	"net/http"
)

func B_Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := md5.Sum([]byte(r.FormValue("password")))
	result := model.SelectWhere(model.Admin{}, "email =? and password =?", "email,password,lastTime,lastIpAddress", email, fmt.Sprintf("%x", password[:]))
	token := r.FormValue("code")
	code := r.FormValue("validate_code")
	if len((*(result.(*[]model.Admin)))) > 0 && emailValidate.GetTokenValue(token) == code {
		cookie := http.Cookie{
			Name:   "token",
			Value:  SessionManager.CreateSession(),
			Path:   "/",
			MaxAge: 3600,
		}
		cookie2 :=
			http.Cookie{
				Name:   "email",
				Value:  email,
				Path:   "/",
				MaxAge: 3600,
			}
		http.SetCookie(w, &cookie)
		http.SetCookie(w, &cookie2)
		http.Redirect(w, r, "/admin/home", http.StatusFound)
		return
	} else {
		t, err := template.ParseFiles("view/backup/login.html")
		if err != nil {
			logger.ErrLog.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		token := emailValidate.GetToken()
		valus := make(map[string]interface{})
		valus["token"] = token
		valus["display"] = "block"
		valus["errorInfo"] = "账号密码或验证码错误"
		t.Execute(w, valus)
	}
}
