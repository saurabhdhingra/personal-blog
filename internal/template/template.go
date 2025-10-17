package template

import (
	"fmt"
	"html/template"
	"net/http"
)

// Base layout template
const LayoutTemplate = `
{{define "layout"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} | Personal Blog</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif; margin: 0; background-color: #f4f7f6; color: #333; }
        header { background-color: #2c3e50; color: white; padding: 1.5rem 0; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .container { max-width: 900px; margin: 0 auto; padding: 0 1.5rem; }
        nav a { color: white; margin-left: 1rem; text-decoration: none; font-weight: 500; transition: color 0.2s; }
        nav a:hover { color: #ecf0f1; text-decoration: underline; }
        main { padding: 2rem 0; }
        .card { background: white; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.05); padding: 1.5rem; margin-bottom: 1.5rem; }
        h1, h2, h3 { color: #2c3e50; margin-top: 0; }
        h1 { font-size: 2.2rem; margin-bottom: 1rem; }
        h2 { font-size: 1.8rem; border-bottom: 2px solid #bdc3c7; padding-bottom: 0.5rem; margin-bottom: 1rem; }
        .article-list-item { border-bottom: 1px solid #eee; padding-bottom: 1rem; margin-bottom: 1rem; }
        .article-list-item:last-child { border-bottom: none; margin-bottom: 0; }
        .article-list-item h3 a { color: #3498db; text-decoration: none; font-size: 1.5rem; }
        .article-list-item p.meta { color: #7f8c8d; font-size: 0.9rem; margin-top: 0.25rem; }
        .btn { display: inline-block; padding: 0.5rem 1rem; border-radius: 4px; text-decoration: none; font-weight: 600; cursor: pointer; transition: background-color 0.2s; }
        .btn-primary { background-color: #3498db; color: white; border: none; }
        .btn-primary:hover { background-color: #2980b9; }
        .btn-danger { background-color: #e74c3c; color: white; border: none; }
        .btn-danger:hover { background-color: #c0392b; }
        .btn-secondary { background-color: #95a5a6; color: white; border: none; }
        .btn-secondary:hover { background-color: #7f8c8d; }
        .form-group { margin-bottom: 1rem; }
        .form-group label { display: block; margin-bottom: 0.5rem; font-weight: 600; color: #555; }
        .form-group input[type="text"], .form-group input[type="password"], .form-group textarea {
            width: 100%;
            padding: 0.75rem;
            border: 1px solid #bdc3c7;
            border-radius: 4px;
            box-sizing: border-box;
        }
        .form-group textarea { min-height: 250px; resize: vertical; }
        .actions { margin-top: 1rem; }
        .actions a { margin-right: 1rem; }
        .login-card { max-width: 400px; margin: 4rem auto; }
    </style>
</head>
<body>
    <header>
        <div class="container">
            <nav style="display: flex; justify-content: space-between; align-items: center;">
                <a href="/"><h1>My Go Blog</h1></a>
                <div>
                    <a href="/">Home</a>
                    <a href="/admin/dashboard">Admin</a>
                </div>
            </nav>
        </div>
    </header>
    <main>
        <div class="container">
            {{template "content" .}}
        </div>
    </main>
</body>
</html>
{{end}}
`

// Guest Home Page
const HomeTemplate = `
{{define "content"}}
    <h2>Latest Articles</h2>
    {{range .Articles}}
        <div class="article-list-item">
            <h3><a href="/article/{{.ID}}">{{.Title}}</a></h3>
            <p class="meta">Published on: {{.PublishedDate}}</p>
        </div>
    {{else}}
        <div class="card">
            <p>No articles published yet.</p>
        </div>
    {{end}}
{{end}}
`

// Guest Article Page
const ArticleTemplate = `
{{define "content"}}
    <div class="card">
        <h1>{{.Article.Title}}</h1>
        <p class="meta">Published on: {{.Article.PublishedDate}}</p>
        <hr style="margin: 1.5rem 0; border: 0; border-top: 1px solid #eee;">
        <div class="article-content">
            {{.Article.Content}}
        </div>
    </div>
    <div class="actions">
        <a href="/" class="btn btn-secondary">← Back to Home</a>
    </div>
{{end}}
`

// Admin Login Page
const LoginTemplate = `
{{define "content"}}
    <div class="card login-card">
        <h2>Admin Login</h2>
        <form method="POST" action="/admin/login">
            <div class="form-group">
                <label for="username">Username</label>
                <input type="text" id="username" name="username" required>
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" id="password" name="password" required>
            </div>
            <button type="submit" class="btn btn-primary" style="width: 100%;">Log In</button>
        </form>
    </div>
{{end}}
`

// Admin Dashboard Page
const DashboardTemplate = `
{{define "content"}}
    <div style="display: flex; justify-content: space-between; align-items: center;">
        <h2>Admin Dashboard</h2>
        <div style="display: flex; gap: 10px;">
            <a href="/admin/add" class="btn btn-primary">Add New Article</a>
            <a href="/admin/logout" class="btn btn-secondary">Logout</a>
        </div>
    </div>
    
    <div class="card" style="margin-top: 1.5rem;">
        <table style="width: 100%; border-collapse: collapse;">
            <thead>
                <tr style="border-bottom: 2px solid #eee;">
                    <th style="text-align: left; padding: 0.75rem 0;">Title</th>
                    <th style="width: 150px; text-align: left; padding: 0.75rem 0;">Date</th>
                    <th style="width: 150px; text-align: right; padding: 0.75rem 0;">Actions</th>
                </tr>
            </thead>
            <tbody>
                {{range .Articles}}
                    <tr style="border-bottom: 1px solid #eee;">
                        <td style="padding: 0.75rem 0; max-width: 600px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">{{.Title}}</td>
                        <td style="padding: 0.75rem 0;">{{.PublishedDate}}</td>
                        <td style="padding: 0.75rem 0; text-align: right;">
                            <a href="/admin/edit/{{.ID}}" class="btn btn-primary" style="padding: 0.4rem 0.6rem; font-size: 0.8rem;">Edit</a>
                            <form action="/admin/delete/{{.ID}}" method="POST" style="display: inline-block;">
                                <button type="submit" class="btn btn-danger" style="padding: 0.4rem 0.6rem; font-size: 0.8rem;" onclick="return confirm('Are you sure you want to delete this article?');">Delete</button>
                            </form>
                        </td>
                    </tr>
                {{else}}
                    <tr>
                        <td colspan="3" style="text-align: center; padding: 1rem;">No articles found. Start writing!</td>
                    </tr>
                {{end}}
            </tbody>
        </table>
    </div>
{{end}}
`

// Admin Article Form (used for both Add and Edit)
const ArticleFormTemplate = `
{{define "content"}}
    <h2>{{.Title}}</h2>
    <div class="card">
        <form method="POST" action="/admin/{{if .Article.ID}}edit/{{.Article.ID}}{{else}}add{{end}}">
            <div class="form-group">
                <label for="title">Title</label>
                <input type="text" id="title" name="title" value="{{.Article.Title}}" required>
            </div>
            <div class="form-group">
                <label for="date">Date of Publication (YYYY-MM-DD)</label>
                <input type="text" id="date" name="date" value="{{if .Article.PublishedDate}}{{.Article.PublishedDate}}{{else}}{{$.CurrentDate}}{{end}}" placeholder="YYYY-MM-DD" required>
            </div>
            <div class="form-group">
                <label for="content">Content</label>
                <textarea id="content" name="content" required>{{.Article.Content}}</textarea>
            </div>
            <div class="actions">
                <button type="submit" class="btn btn-primary">Save Article</button>
                <a href="/admin/dashboard" class="btn btn-secondary">Cancel</a>
            </div>
        </form>
    </div>
{{end}}
`

func ParseAndExecute(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.New("layout").Parse(LayoutTemplate)
	if err != nil {
		http.Error(w, "Internal Server Error: Layout template failed to parse.", http.StatusInternalServerError)
		return
	}
	t, err = t.Parse(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error: Content template failed to parse.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := t.Execute(w, data); err != nil {
		fmt.Printf("Template execution error: %v\n", err)
		http.Error(w, "Internal Server Error during rendering.", http.StatusInternalServerError)
	}
}