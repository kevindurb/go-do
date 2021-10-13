package handlers

import (
	"kevindurb/go-do/utils"
	"net/http"
)

func Status(response http.ResponseWriter, request *http.Request) {
	utils.RespondSuccess(response, "hello!")
}
