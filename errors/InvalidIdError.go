package custom_errors

import "sync"

var invalidIdErrorOnce sync.Once
var invalidIdErrorInstance *InvalidIdError

type InvalidIdError struct {
	message string
}

func NewInvalidIdError() *InvalidIdError {
	invalidIdErrorOnce.Do(func() {
		invalidIdErrorInstance = &InvalidIdError{
			message: "Invalid id",
		}
	})
	return invalidIdErrorInstance
}

func (err *InvalidIdError) Error() string {
	return err.message
}
