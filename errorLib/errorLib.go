package errorlib

import "net/http"

// Error ..
type Error struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"errorCode"`
	Message    string `json:"errorMessage"`
}

const (
	NotFoundCode              string = "NOT_FOUND"
	InternalServerErrorCode   string = "INTERNAL_SERVER_ERROR"
	BadRequestCode            string = "BAD_REQUEST"
	ResourceAlreadyExistsCode string = "RESOURCE_ALREADY_EXISTS"
	UnauthorizedCode          string = "UNAUTHORIZED"
	ForbiddenCode             string = "FORBIDDEN_CODE"
)

func NotFound(message string) *Error {
	return &Error{StatusCode: http.StatusNotFound, Code: NotFoundCode, Message: message}
}

func InternalServerError(message string) *Error {
	return &Error{StatusCode: http.StatusInternalServerError, Code: InternalServerErrorCode, Message: message}
}

func Forbidden(message string) *Error {
	return &Error{StatusCode: http.StatusForbidden, Code: ForbiddenCode, Message: message}
}

func BadRequest(message string) *Error {
	return &Error{StatusCode: http.StatusBadRequest, Code: BadRequestCode, Message: message}
}

func ResourceAlreadyExist(message string) *Error {
	return &Error{StatusCode: http.StatusConflict, Code: ResourceAlreadyExistsCode, Message: message}
}

func Unauthorized(message string) *Error {
	return &Error{StatusCode: http.StatusUnauthorized, Code: UnauthorizedCode, Message: message}
}
