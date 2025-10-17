package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"personal-blog/internal/model"
	"personal-blog/internal/storage"
	"personal-blog/internal/template"
)

// HandleDashboard displays the list of articles for editing/deleting.
func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	articles, err := storage.LoadAllArticles()
	if err != nil && !os.IsNotExist(err) && len(articles) == 0 {
		fmt.Printf("Error loading articles for dashboard: %v\n", err)
	}

	data := struct {
		Title    string
		Articles []model.Article
	}{
		Title:    "Admin Dashboard",
		Articles: articles,
	}

	template.ParseAndExecute(w, template.DashboardTemplate, data)
}

// HandleAddArticle displays the form for adding a new article or processes the submission.
func HandleAddArticle(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title       string
		Article     model.Article
		CurrentDate string
	}{
		Title:       "Add New Article",
		CurrentDate: time.Now().Format("2006-01-02"),
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		title := r.FormValue("title")
		content := r.FormValue("content")
		dateStr := r.FormValue("date")

		// ID is intentionally empty here; storage.SaveArticle will generate it.
		newArticle := model.Article{
			Title:         title,
			Content:       content,
			PublishedDate: dateStr,
		}

		if err := storage.SaveArticle(newArticle); err != nil {
			fmt.Printf("Error saving article: %v\n", err)
			http.Error(w, "Error saving article to file.", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	template.ParseAndExecute(w, template.ArticleFormTemplate, data)
}

// HandleEditArticle displays the form for editing an article or processes the update.
func HandleEditArticle(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.NotFound(w, r)
		return
	}
	id := parts[3]

	article, err := storage.LoadArticle(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Title       string
		Article     model.Article
		CurrentDate string
	}{
		Title:       "Edit Article: " + article.Title,
		Article:     article,
		CurrentDate: article.PublishedDate,
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		article.Title = r.FormValue("title")
		article.Content = r.FormValue("content")
		article.PublishedDate = r.FormValue("date")
		// article.ID is already set from the loaded article

		if err := storage.SaveArticle(article); err != nil {
			fmt.Printf("Error updating article: %v\n", err)
			http.Error(w, "Error updating article file.", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	template.ParseAndExecute(w, template.ArticleFormTemplate, data)
}

// HandleDeleteArticle processes the request to delete an article.
func HandleDeleteArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.NotFound(w, r)
		return
	}
	id := parts[3]

	if err := storage.DeleteArticle(id); err != nil {
		fmt.Printf("Error deleting article: %v\n", err)
		http.Error(w, "Error deleting article file.", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}