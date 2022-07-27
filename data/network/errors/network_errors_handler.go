package network_errors

import (
	"errors"
	"net/http"
	custom_errors "relembrando_2/errors"
	"strconv"
)

func Handle(err error) error {
	if err == nil {
		return nil
	} else if custom_errors.IsACustomError(err) {
		return err
	} else if errors.Is(err, strconv.ErrSyntax) ||
		errors.Is(err, strconv.ErrRange) || errors.Is(err, &strconv.NumError{}) {
		return custom_errors.NewInvalidIdError()
	} else {
		return custom_errors.NewInternalError()
	}
}

func HandleHttpStatus(statusCode int) error {
	if statusCode == http.StatusNotFound {
		return custom_errors.NewNotFoundError()
	} else {
		return nil
	}
}
