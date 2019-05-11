package models

type Books struct {
	ID          int    `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Author      string `form:"author" json:"author"`
	Description string `form:"description" json:"description"`
	CreatedAt   string `form:"created_at" json:"created_at"`
	UpdatedAt   string `form:"updated_at" json:"updated_at"`
}
