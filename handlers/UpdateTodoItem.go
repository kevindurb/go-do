package handlers

import (
	"encoding/json"
	"kevindurb/go-do/models"
	"kevindurb/go-do/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateTodoItem(response http.ResponseWriter, request *http.Request) {
	db := utils.GetConnection()
	id, _ := strconv.ParseUint(mux.Vars(request)["id"], 10, 64)

	var todoItem models.TodoItem
	utils.ParseBody(request, &todoItem)
	db.Model(&models.TodoItem{ID: uint(id)}).Updates(&todoItem)
	db.First(&todoItem, uint(id))
	json.NewEncoder(response).Encode(todoItem)
}