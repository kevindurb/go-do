package handlers

import (
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
	"net/http"
)

func CreateTodoItem(response http.ResponseWriter, request *http.Request) {
	db := utils.GetConnection()

	var todoItem models.TodoItem
	utils.ParseBody(request, &todoItem)
	db.Create(&todoItem)
	response.WriteHeader(http.StatusCreated)
}