package handlers

import (
	"encoding/json"
	"io/ioutil"
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
	"net/http"
)

func CreateTodoItem(response http.ResponseWriter, request *http.Request) {
	db := utils.GetConnection()
	defer request.Body.Close()
  body, _ := ioutil.ReadAll(request.Body)

	var todoItem models.TodoItem
	json.Unmarshal(body, &todoItem)

	db.Create(&todoItem)
}