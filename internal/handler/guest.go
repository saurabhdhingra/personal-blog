package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"personal-blog/internal/model"
	"personal-blog/internal/storage"
	"personal-blog/internal/template"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	articles, err := storage.LoadAllArticles()
	if err != nil && !os.IsNotExist(err) && len(articles) == 0 {
		fmt.Printf("Error loading articles: %v\n", err)
	}

	data := struct {
		Title    string
		Articles []model.Article
	}{
		Title:    "Home",
		Articles: articles,
	}

	template.ParseAndExecute(w, template.HomeTemplate, data)
}

func HandleArticle(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.NotFound(w, r)
		return
	}
	id := parts[2]

	article, err := storage.LoadArticle(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Title   string
		Article model.Article
	}{
		Title:   article.Title,
		Article: article,
	}

	template.ParseAndExecute(w, template.ArticleTemplate, data)
}
