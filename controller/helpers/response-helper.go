package helpers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, object interface{}, statusCode int) error {
	if object == nil {
		w.WriteHeader(statusCode)
		return nil
	}

	if err := json.NewEncoder(w).Encode(object); err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	return nil
}

func Redirect(w http.ResponseWriter, r *http.Request, url string) error {
	http.Redirect(w, r, url, http.StatusSeeOther)
	return nil
}

func SendResponseStatusOk(w http.ResponseWriter, object interface{}) error {
	return SendResponse(w, object, 200)
}
