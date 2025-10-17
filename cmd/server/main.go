package main

import (
	"fmt"
	"net/http"

	"personal-blog/internal/auth"
	"personal-blog/internal/handler"
	"personal-blog/internal/model"
	"personal-blog/internal/storage"
)

func main(){
	storage.InitFS()
	auth.InitSessionKey()

	http.HandleFunc("/", handler.HandleHome)
	http.HandleFunc("/article/", handler.HandleArticle)
	
	// Dashboard
	http.HandleFunc("/admin/dashboard", auth.RequireAuth(handler.HandleDashboard))
	// Add Article (GET for form, POST for submission)
	http.HandleFunc("/admin/add", auth.RequireAuth(handler.HandleAddArticle))
	// Edit Article (GET for form, POST for submission)
	http.HandleFunc("/admin/edit/", auth.RequireAuth(handler.HandleEditArticle))
	// Delete Article (POST request)
	http.HandleFunc("/admin/delete/", auth.RequireAuth(handler.HandleDeleteArticle))

	http.HandleFunc("/admin/login", handler.HandleLogin)
	http.HandleFunc("/admin/logout", handler.HandleLogout)

	port := ":8080"
	fmt.Printf("Starting blog server on http://localhost%s\n", port)
	fmt.Printf("Admin Login: %s / %s\n", model.AdminCreds.Username, model.AdminCreds.Password)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}