package handler

import (
	"net/http"

	"personal-blog/internal/auth"
	"personal-blog/internal/model"
	"personal-blog/internal/template"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if auth.IsAuthenticated(r) {
		http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == model.AdminCreds.Username && password == model.AdminCreds.Password {
			auth.SetAuthCookie(w)
			http.Redirect(w, r, "/admin/dashboard", http.StatusFound)
			return
		}

		http.Redirect(w, r, "/admin/login?error=true", http.StatusFound)
		return
	}

	template.ParseAndExecute(w, template.LoginTemplate, struct{ Title string }{Title: "Admin Login"})
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	auth.ClearAuthCookie(w)
	http.Redirect(w, r, "/", http.StatusFound)
}