package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

func GetObjectFromJson[T any](r *http.Request, object *T) ( error) {
	if err := json.NewDecoder(r.Body).Decode(&object); err != nil {
		return err
	}

    if err := validator.New().Struct(object); err != nil {
        // Validation failed, handle the error
        errs := err.(validator.ValidationErrors)
        return errors.New(fmt.Sprintf("Validation error: %s", errs)) 
    }

	return  nil
}