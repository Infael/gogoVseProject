package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Infael/gogoVseProject/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func GetObjectFromJson[T any](r *http.Request, object *T) error {
	if err := json.NewDecoder(r.Body).Decode(&object); err != nil {
		return err
	}

	if err := validator.New().Struct(object); err != nil {
		// Validation failed, handle the error
		errs := err.(validator.ValidationErrors)
		return errors.New(fmt.Sprintf("validation error: %s", errs))
	}

	return nil
}

func GetIdFromRequest(r *http.Request) (*uint64, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, utils.ErrorBadRequest(errors.New("id is required"))
	}
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, utils.ErrorBadRequest(errors.New("id must be a number"))
	}
	return &parsedId, nil
}
