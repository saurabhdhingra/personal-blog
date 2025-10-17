package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"personal-blog/internal/model"
)

const articlesDir = "articles"

func InitFS() {
	if _, err := os.Stat(articlesDir); os.IsNotExist(err) {
		if err := os.Mkdir(articlesDir, 0755); err != nil {
			fmt.Printf("Error creating articles directory: %v\n", err)
			os.Exit(1)
		}
	}
}

func slugify(s string) string {
	s = strings.ToLower(s)
	reg, _ := regexp.Compile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

func generateID(title string) string {
	slug := slugify(title)

	if _, err := os.Stat(filepath.Join(articlesDir, slug+".json")); err == nil {
		slug = fmt.Sprintf("%s-%d", slug, time.Now().Unix())
	}
	return slug
}

func SaveArticle(article model.Article) error {
	if article.ID == "" {
		article.ID = generateID(article.Title)
	}

	filePath := filepath.Join(articlesDir, article.ID+".json")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(article)
}

func LoadArticle(id string) (model.Article, error) {
	filePath := filepath.Join(articlesDir, id+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return model.Article{}, err
	}
	var article model.Article
	if err := json.Unmarshal(data, &article); err != nil {
		return model.Article{}, err
	}
	return article, nil
}

func LoadAllArticles() ([]model.Article, error) {
	var articles []model.Article
	files, err := os.ReadDir(articlesDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			id := strings.TrimSuffix(file.Name(), ".json")
			article, err := LoadArticle(id)
			if err == nil {
				articles = append(articles, article)
			}
		}
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].PublishedDate > articles[j].PublishedDate
	})

	return articles, nil
}

func DeleteArticle(id string) error {
	filePath := filepath.Join(articlesDir, id+".json")
	return os.Remove(filePath)
}