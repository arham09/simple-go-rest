package models

import (
	"book-rest/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Books struct {
	ID          int    `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Author      string `form:"author" json:"author"`
	Description string `form:"description" json:"description"`
	Status      int    `form:"status" json:"status"`
	CreatedAt   string `form:"created_at" json:"created_at"`
	UpdatedAt   string `form:"updated_at" json:"updated_at"`
}

func GetBooks() (*[]Books, error) {
	var book Books
	var books []Books

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, author, description, status, created_at, updated_at FROM books WHERE status = 1")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.Status, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Fatal(err.Error())

		} else {
			books = append(books, book)
		}
	}

	return &books, err
}

func GetBook(bookId int) (*Books, error) {
	var book Books

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, author, description, status, created_at, updated_at FROM books WHERE status = 1 AND id = ?", bookId)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.Status, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Fatal(err.Error())
		}
	}

	return &book, err
}
