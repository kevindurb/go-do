package main

import (
	"kevindurb/go-do/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/status", handlers.Status)
	router.HandleFunc("/items", handlers.CreateTodoItem).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
