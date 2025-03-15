package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type HttpError struct {
    Code int `json:"code"`
    Message string `json:"message"`
}

func NewHttpError(code int, message string) HttpError{
    return HttpError{
        Code: code,
        Message: message,
    }
}

func (e HttpError) Write(w http.ResponseWriter) {
    w.WriteHeader(e.Code)

    r, err := json.Marshal(e)
    if err != nil {
        w.Write([]byte ("{'message':\"Internal Server Error\", 'code':-1}"))
    }

    w.Write(r)
}

func getValidationErrors(ve validator.ValidationErrors) map[string]string {
	errors := make(map[string]string, len(ve))
	for _, fe := range ve {
		errors[fe.Field()] = fmt.Sprintf("%s",fe.Error())
	}
	return errors
}
