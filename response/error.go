package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json: "status"`
	Message string `json: "mesage"`
}

// GetData implements Response.
func (e *ErrorResponse) GetData() interface{} {
	panic("unimplemented")
}

func InternalServerError(msg string) Response {
	return errorR(msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return errorR(msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return errorR(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return errorR(msg, http.StatusForbidden)
}

func BadRequest(msg string) Response {
	return errorR(msg, http.StatusBadRequest)
}

func errorR(msg string, code int) Response {
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
