package transactions

import (
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
)

func CreateTodoItem(todoItem *models.TodoItem) {
	db := utils.GetConnection()

	db.Create(todoItem)
}