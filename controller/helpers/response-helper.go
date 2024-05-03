package helpers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, object interface{}, statusCode int) error {
	if err := json.NewEncoder(w).Encode(object); err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	return nil
}

func SendResponseStatusOk(w http.ResponseWriter, object interface{}) error {
	return SendResponse(w, object, 200)
}
