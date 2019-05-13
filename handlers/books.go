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
	var response models.ResponseBooks

	books, err := models.GetBooks()
	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Data = *books

	middlewares.Response(w, http.StatusOK, response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	var response models.ResponseBook

	vars := mux.Vars(r)
	bookId := vars["bookId"]

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
	status := 1
	createdAt := time.Now()
	updatedAt := time.Now()

	_, err = db.Exec("INSERT INTO books(name, author, description, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", name, author, description, status, createdAt, updatedAt)

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

	_, err = db.Exec("UPDATE books SET name = ?, author = ?, description = ?, updated_at = ? WHERE id = ?", name, author, description, updatedAt, bookId)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "Successfully Updated"

	middlewares.Response(w, http.StatusOK, response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	db := config.Connect()
	defer db.Close()

	status := 0
	updatedAt := time.Now()

	_, err := db.Exec("UPDATE books SET status = ?, updated_at = ? WHERE id = ?", status, updatedAt, bookId)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "Successfully Deleted"

	middlewares.Response(w, http.StatusOK, response)
}
