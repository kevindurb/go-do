package handlers

import (
	"encoding/json"
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
	"net/http"
)

func UpdateTodoItem(response http.ResponseWriter, request *http.Request) {
	db := utils.GetConnection()

	var todoItem models.TodoItem
	utils.ParseBody(request, &todoItem)
	db.Updates(&todoItem)
	db.First(&todoItem, todoItem.ID)
	json.NewEncoder(response).Encode(todoItem)
}