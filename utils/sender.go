package utils

import (
	"encoding/json"
	"net/http"
)

//ResponseSender function
func ResponseSender(res http.ResponseWriter, req *http.Request, data interface{}) {
	req.Header.Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}
