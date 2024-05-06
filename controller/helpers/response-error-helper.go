package helpers

import (
	"log"
	"net/http"

	"github.com/Infael/gogoVseProject/utils"
)

func SendError(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {
	case utils.StatusError:
		// We can retrieve the status here and write out a specific
		// HTTP status code.
		log.Printf("Request %s - Caller %s - Responded with: %d - %s", r.URL, e.Caller, e.StatusCode(), e.Error())

		http.Error(w, http.StatusText(e.StatusCode()), e.Code)
		return
	default:
		// Any error types we don't specifically look out for default
		// to serving a HTTP Internal Server Error
		log.Printf("Unspecified error: %d - %s", http.StatusInternalServerError, e.Error())

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
