package handlers

import (
	"kevindurb/go-do/models"
	"kevindurb/go-do/transactions"
	"kevindurb/go-do/utils"
	"net/http"
)

type TodoItemUpdateBody struct {
	Description string `json:"description"`
	Done bool `json:"done"`
}

func UpdateTodoItem(response http.ResponseWriter, request *http.Request) {
	id := utils.ParseUIntParam(request, "id")

	var updates TodoItemUpdateBody
	utils.ParseBody(request, &updates)

	todoItem := models.TodoItem{
		ID: id,
		Description: updates.Description,
		Done: updates.Done,
	}

	transactions.UpdateTodoItem(&todoItem)
	utils.RespondSuccess(response, todoItem)
}