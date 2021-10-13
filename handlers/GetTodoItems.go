package handlers

import (
	"kevindurb/go-do/models"
	"kevindurb/go-do/transactions"
	"kevindurb/go-do/utils"
	"net/http"
)

func GetTodoItems(response http.ResponseWriter, request *http.Request) {
	var items []models.TodoItem
	transactions.GetTodoItems(&items)
	utils.RespondSuccess(response, items)
}