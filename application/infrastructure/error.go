package infrastructure

import "errors"

var (
	// http error 500
	InternalServerError = errors.New("Internal Server Error")
	// http error 404
	NotFound = errors.New("Resource Not Found")
	// http error 402
	BadRequest = errors.New("Given Params is Not Valid")
)
