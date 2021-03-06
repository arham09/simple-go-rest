package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/arham09/book-rest/handlers"
	"github.com/arham09/book-rest/middlewares"
)

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/books", middlewares.Authorized(handlers.GetBooks)).Methods("GET")
	api.HandleFunc("/books/{bookId}", middlewares.Authorized(handlers.GetBook)).Methods("GET")
	api.HandleFunc("/books/add", middlewares.Authorized(handlers.CreateBook)).Methods("POST")
	api.HandleFunc("/books/edit/{bookId}", middlewares.Authorized(handlers.UpdateBook)).Methods("PUT")
	api.HandleFunc("/books/delete/{bookId}", middlewares.Authorized(handlers.DeleteBook)).Methods("DELETE")

	api.HandleFunc("/users/register", handlers.RegisterUser).Methods("POST")
	api.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	fmt.Println("Connected to port 2019")
	http.Handle("/", middlewares.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, router)))
	log.Fatal(http.ListenAndServe(":2019", nil))
}
