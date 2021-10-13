package transactions

import (
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
)

func UpdateTodoItem(todoItem *models.TodoItem) {
	db := utils.GetConnection()
	db.Model(todoItem).Updates(todoItem)
	db.First(todoItem, todoItem.ID)
}