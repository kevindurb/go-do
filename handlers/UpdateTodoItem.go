package handlers

import (
	"encoding/json"
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
	"net/http"
)

func UpdateTodoItem(response http.ResponseWriter, request *http.Request) {
	db := utils.GetConnection()
	id := utils.ParseUIntParam(request, "id")

	var todoItem models.TodoItem
	utils.ParseBody(request, &todoItem)
	db.Model(&models.TodoItem{ID: id}).Updates(&todoItem)
	db.First(&todoItem, id)
	json.NewEncoder(response).Encode(todoItem)
}