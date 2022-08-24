package common

import (
	"encoding/json"
	"net/http"
)

// ResponseJson
func ResponseJson(w http.ResponseWriter, statusCode int, respStruct interface{}) {
	b, err := json.Marshal(&respStruct)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(b)
}
