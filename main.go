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
	router.HandleFunc("/items", handlers.GetTodoItems).Methods(http.MethodGet)
	router.HandleFunc("/items", handlers.UpdateTodoItem).Methods(http.MethodPut)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
