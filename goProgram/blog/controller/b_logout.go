package controller

import (
	"net/http"
)

func B_Logout(w http.ResponseWriter, r *http.Request) {
	token := SessionManager.CookieIsExist(r, "token")
	SessionManager.DestroySession(token)
	http.Redirect(w, r, "/admin", http.StatusFound)
	return
}
