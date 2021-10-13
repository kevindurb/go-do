package transactions

import (
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
)

func GetTodoItems(items *[]models.TodoItem) {
	db := utils.GetConnection()

	db.Find(items)
}