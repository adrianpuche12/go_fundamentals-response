package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json: "status"`
	Message string `json: "mesage"`
}

func InternalServerError(msg string) Response {
	return error(msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return error(msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return error(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return error(msg, http.StatusForbidden)
}

func BadRequest(msg string) Response {
	return error(msg, http.StatusBadRequest)
}

func error(msg string, code int) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  code,
	}
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func (e ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (s *ErrorResponse) Getdata() interface{} {
	return nil
}
