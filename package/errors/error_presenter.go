package errors

import (
	"fmt"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func (err ErrorResponse) Error() string {
	return err.Message
}

type ErrorLog struct {
	Code     int
	Message  string
	ErrorLog error
}

func (err *ErrorLog) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *ErrorLog) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *ErrorLog) Error() string {
	return fmt.Sprintf("Error - code: %d, message: %s, error: %s", err.Code, err.Message, err.ErrorLog)
}

func New(errResp *ErrorResponse, err error) *ErrorLog {
	return &ErrorLog{Code: errResp.Code, Message: errResp.Message, ErrorLog: err}
}

func IsUserNotFoundError(err error) bool {
	code, _ := DecodeError(err)

	return code == UserNotFoundError.Code
}

func DecodeError(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *ErrorLog:
		return typed.Code, typed.Message
	case *ErrorResponse:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
