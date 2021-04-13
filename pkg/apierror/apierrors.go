package apierrors

import (
	"fmt"
)

type CauseList []interface{}

type ApiError interface {
	Message() string
	Code() string
	Status() int
	Error() string
}

type apiErr struct {
	ErrorMessage string `json:"message"`
	ErrorCode    string `json:"error"`
	ErrorStatus  int    `json:"status"`
}

func (c CauseList) ToString() string {
	return fmt.Sprint(c)
}

func (e apiErr) Code() string {
	return e.ErrorCode
}

func (e apiErr) Error() string {
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d", e.ErrorMessage, e.ErrorCode, e.ErrorStatus)
}

func (e apiErr) Status() int {
	return e.ErrorStatus
}

func (e apiErr) Message() string {
	return e.ErrorMessage
}

func NewApiError(message string, error string, status int) ApiError {
	return apiErr{message, error, status}
}
