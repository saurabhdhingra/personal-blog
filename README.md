# Simple Go Blog Server

This project is a personal blog application built entirely with Go, following a modular architecture. It implements server-side rendering, file-system-based storage (JSON files), and basic cookie-based session authentication for the administrative area.

## Architecture

The application is structured into logical packages to separate concerns (Model, Storage, Authentication, Handlers, and Templates), promoting maintainability and adherence to Go best practices.

```
blog-server/
├── cmd/
│   └── server/
│       └── main.go       <-- Server Entry Point & Router
├── internal/
│   ├── auth/             <-- Authentication (Session management, Middleware)
│   ├── handler/          <-- HTTP Request Handlers (Guest, Auth, Admin CRUD)
│   ├── model/            <-- Data Structures (Article, Credentials)
│   ├── storage/          <-- Persistence Logic (Read/Write JSON files on filesystem)
│   └── template/         <-- HTML Templates and Rendering Utilities
```

## Getting Started

Prerequisites

You need Go (version 1.18 or higher) installed on your system.

Setup and Running

Initialize the Go Module:

```
go mod init blog-server
```

Run the Server:
The server will automatically create an ./articles directory if it doesn't already exist.

```
go run cmd/server/main.go
```

The application will start on port 8080.

## Usage

### Guest Section (Public Access)

These pages are accessible by anyone visiting the blog.

URL             Description

/               Home Page: Displays a list of all published articles, sorted by publication date.

/article/{id}   Article Page: Shows the full content and details of a specific article, identified by its slug/ID.

### Admin Section (Authenticated Access)

These pages require login via cookie-based authentication.

```
URL                     Description

/admin/login            Displays the login form and processes the credentials.

/admin/dashboard        Dashboard: Lists all articles with options to Edit or Delete.

/admin/add              Displays the form for creating a new article.

/admin/edit/{id}        Displays the form pre-filled with the data of an existing article for modification.

/admin/delete/{id}      Processes the POST request to remove an article file.

/admin/logout           Clears the admin session cookie and redirects to the home page.
```

## Admin Credentials (Hardcoded)

For this initial implementation, authentication uses hardcoded values.

```
Field                           Value

Username                        admin

Password                        password123
```

Navigate to http://localhost:8080/admin/login to access the dashboard.

## Data Storage

Articles are stored as individual JSON files within the ./articles directory. Each file contains the article's title, content, unique ID (derived from the title), and publication date.

```
// Example: articles/hello-world.json
{
  "id": "hello-world",
  "title": "Hello World: My First Post",
  "content": "This is the content of my very first blog post...",
  "published_date": "2024-01-01"
}
```

## Acknowledgement
https://roadmap.sh/projects/personal-blog