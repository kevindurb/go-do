package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, value interface{}) {
	defer request.Body.Close()
  body, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, value)
}