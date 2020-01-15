package misc

import (
	"net/http"

	_error "product-microservice/application/infrastructure"

	"github.com/sirupsen/logrus"
)

func GetErrorStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case _error.InternalServerError:
		return http.StatusInternalServerError
	case _error.NotFound:
		return http.StatusNotFound
	case _error.BadRequest:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
