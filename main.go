package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"book-rest/handlers"
	"book-rest/middlewares"
)

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	api.HandleFunc("/books/{bookId}", handlers.GetBook).Methods("GET")

	fmt.Println("Connected to port 2019")
	http.Handle("/", middlewares.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, router)))
	log.Fatal(http.ListenAndServe(":2019", nil))
}
