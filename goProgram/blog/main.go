package main

import (
	"blog/controller"
	"blog/logger"
	"blog/model"
	_ "blog/model"
	"html/template"
	"net/http"
)

func main() {
	logger.InfoLog.Println("Running....")
	logger.ErrLog.Println("Running....")
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t, err := template.ParseFiles("view/index/404.html")
			if err != nil {
				logger.ErrLog.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			t.Execute(w, nil)
		} else {
			controller.Home(w, r)
		}
	})
	//前台页面
	mux.HandleFunc("/about", controller.About)
	mux.HandleFunc("/article", controller.Atricle)
	mux.HandleFunc("/saveComment", controller.SaveComment)
	mux.HandleFunc("/comment", controller.Comment)
	mux.HandleFunc("/friend", controller.Friend)
	mux.HandleFunc("/search", controller.Search)
	mux.HandleFunc("/tag/", controller.Tag)
	mux.HandleFunc("/resume", controller.Resume)
	//后台页面
	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/admin" {
			t, err := template.ParseFiles("view/index/404.html")
			if err != nil {
				logger.ErrLog.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			t.Execute(w, nil)
		} else {
			controller.B_Index(w, r)
		}
	})
	mux.HandleFunc("/admin/sendEmail", controller.B_SendEmail)
	mux.HandleFunc("/admin/login", controller.B_Login)
	mux.HandleFunc("/admin/home", controller.B_Home)
	mux.HandleFunc("/admin/logout", controller.B_Logout)
	mux.HandleFunc("/admin/article", controller.B_Article)
	mux.HandleFunc("/admin/article/delete", controller.B_ArticleDelete)
	mux.HandleFunc("/admin/article/save", controller.B_SaveArticle)
	mux.HandleFunc("/admin/article/view", controller.B_ArticleView)
	mux.HandleFunc("/admin/article/acdelete", controller.B_DeleteArticleComment)
	mux.HandleFunc("/admin/comment", controller.B_Comment)
	mux.HandleFunc("/admin/article/cdelete", controller.B_DeleteComment)
	defer logger.ErrLog.Println("Stopped....")
	defer logger.InfoLog.Println("Stopped....")
	defer model.Close()
	go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusFound)
	}))
	http.ListenAndServeTLS(":443", "", "", mux)

}
