package model

type Article struct {
	ID				string	`json:"id"`
	Title			string	`json:"title"`
	Content			string	`json:"content"`
	PublishedDate	string	`json:"published_date"`
}

type AdminCredentials struct {
	Username		string
	Password	string
}

var AdminCreds = AdminCredentials{
	Username: "admin",
	Password: "password123",
}