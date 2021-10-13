package handlers

import (
	"kevindurb/go-do/models"
	"kevindurb/go-do/transactions"
	"kevindurb/go-do/utils"
	"net/http"
)

type TodoItemCreateBody struct {
	Description string `json:"description"`
}

func CreateTodoItem(response http.ResponseWriter, request *http.Request) {
	var newTodoItem TodoItemCreateBody
	utils.ParseBody(request,  &newTodoItem)

	todoItem := models.TodoItem{
		Description: newTodoItem.Description,
	}
	transactions.CreateTodoItem(&todoItem)
	utils.RespondCreated(response, todoItem)
}