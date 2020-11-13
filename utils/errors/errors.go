package errors

import "net/http"

//RestError : To handle the error
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

//NewBadRequest : To generate the bad request error
func NewBadRequest(message string) *RestError {
	err := &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
	return err
}

//NewNotFound : it returns an error when request obj is not present in the database or not
func NewNotFound(message string) *RestError {
	err := &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
	return err
}
