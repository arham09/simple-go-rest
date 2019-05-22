package handlers

import (
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

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
	var response models.ResponseBook

	vars := mux.Vars(r)
	bookId, err := strconv.Atoi(vars["bookId"])
	if err != nil {
		log.Print(err)
	}

	book, err := models.GetBook(&bookId)
	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Data = *book

	middlewares.Response(w, http.StatusOK, response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		return
	}

	name := r.Form.Get("name")
	author := r.Form.Get("author")
	description := r.Form.Get("description")

	err = models.InsertBook(&name, &author, &description)

	if err != nil {
		log.Print("ada error")
		log.Print(err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Successfully inserted"

	middlewares.Response(w, http.StatusOK, response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	vars := mux.Vars(r)
	bookId, err := strconv.Atoi(vars["bookId"])
	if err != nil {
		log.Print(err)
	}

	err = r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	name := r.Form.Get("name")
	author := r.Form.Get("author")
	description := r.Form.Get("description")

	err = models.EditBook(&bookId, &name, &author, &description)

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
	bookId, err := strconv.Atoi(vars["bookId"])
	if err != nil {
		log.Print(err)
	}

	err = models.RemoveBook(&bookId)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "Successfully Deleted"

	middlewares.Response(w, http.StatusOK, response)
}
