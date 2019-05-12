package handlers

import (
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"book-rest/config"
	"book-rest/middlewares"
	"book-rest/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	var books []models.Books
	var response models.ResponseBooks

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, author, description, created_at, updated_at FROM books")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Fatal(err.Error())

		} else {
			books = append(books, book)
		}
	}

	response.Status = http.StatusOK
	response.Data = books

	middlewares.Response(w, http.StatusOK, response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	var response models.ResponseBook

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, author, description, created_at, updated_at FROM books WHERE id = ?", bookId)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Description, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Fatal(err.Error())
		}
	}

	response.Status = http.StatusOK
	response.Data = book

	middlewares.Response(w, http.StatusOK, response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	name := r.Form.Get("name")
	author := r.Form.Get("author")
	description := r.Form.Get("description")
	createdAt := time.Now()
	updatedAt := time.Now()

	_, err = db.Exec("INSERT INTO books(name, author, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", name, author, description, createdAt, updatedAt)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "Successfully inserted"

	middlewares.Response(w, http.StatusOK, response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	db := config.Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	name := r.Form.Get("name")
	author := r.Form.Get("author")
	description := r.Form.Get("description")
	updatedAt := time.Now()

	_, err = db.Exec("UPDATE books set name = ?, author = ?, description = ?, updated_at = ? where id = ?", name, author, description, updatedAt, bookId)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "Successfully Updated"

	middlewares.Response(w, http.StatusOK, response)
}
