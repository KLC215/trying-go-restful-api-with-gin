package errors

var (
	// Common errors
	OK                  = &ErrorResponse{Code: 0, Message: "OK"}
	InternalServerError = &ErrorResponse{Code: 10001, Message: "Internal server error"}
	BindError           = &ErrorResponse{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// User errors
	UserNotFoundError = &ErrorResponse{Code: 20102, Message: "User does not found!"}
)
