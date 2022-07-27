package custom_errors

import (
	"errors"
)

func IsACustomError(err error) bool {
	return errors.Is(err, NewInvalidIdError()) ||
		errors.Is(err, NewNotFoundError()) || errors.Is(err, NewInternalError())
}
