package handlers

import (
	"encoding/json"
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
	"net/http"
)

func GetTodoItems(response http.ResponseWriter, request *http.Request) {
	db := utils.GetConnection()

	var items []models.TodoItem

	db.Find(&items)

	json.NewEncoder(response).Encode(items)
}