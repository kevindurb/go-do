package handlers

import (
	"encoding/json"
	"net/http"
)

func Status(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("Hello World")
}
