package custom_errors

import "sync"

var internalErrorOnce sync.Once
var internalErrorInstance *InternalError

type InternalError struct {
	message string
}

func NewInternalError() *InternalError {
	internalErrorOnce.Do(func() {
		internalErrorInstance = &InternalError{
			message: "Internal error",
		}
	})
	return internalErrorInstance
}

func (err *InternalError) Error() string {
	return err.message
}
