package errs

import (
	"net/http"
)

type AppErr struct {
	Code    int `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func (e AppErr) AsMessage() *AppErr{
	return &AppErr{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Code: http.StatusNotFound,
		Message: message ,
	}
}

func NewUnexpectedError(message string) *AppErr {
	return &AppErr{
		Code: http.StatusInternalServerError,
		Message: message ,
	}
}