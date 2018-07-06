package errors

var (
	// Common errors
	OK                  = &ErrorResponse{Code: 0, Message: "OK"}
	InternalServerError = &ErrorResponse{Code: 10001, Message: "Internal server error"}
	BindError           = &ErrorResponse{Code: 10002, Message: "Error occurred while binding the request body to the struct"}
	ValidationError     = &ErrorResponse{Code: 20001, Message: "Validation failed"}
	DatabaseError       = &ErrorResponse{Code: 20002, Message: "Database error"}
	TokenError          = &ErrorResponse{Code: 20003, Message: "Error occurred when signing JWT"}

	// User errors
	EncryptError           = &ErrorResponse{Code: 20101, Message: "Error occurred when encrypting password"}
	UserNotFoundError      = &ErrorResponse{Code: 20102, Message: "User does not found"}
	TokenInvalidError      = &ErrorResponse{Code: 20103, Message: "Invalid token"}
	PasswordIncorrectError = &ErrorResponse{Code: 20104, Message: "Incorrect password"}
)
