package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(request *http.Request, value interface{}) {
	defer request.Body.Close()
  body, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, value)
}

func ParseUIntParam(request *http.Request, key string) uint {
	vars := mux.Vars(request)
	intVal, _ := strconv.ParseUint(vars[key], 10, 64)
	return uint(intVal)
}

func RespondCreated(response http.ResponseWriter, value interface{}) {
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(value)
}

func RespondSuccess(response http.ResponseWriter, value interface{}) {
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(value)
}