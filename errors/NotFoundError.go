package custom_errors

import "sync"

var notFoundErrorOnce sync.Once
var notFoundErrorInstance *NotFound

type NotFound struct {
	message string
}

func NewNotFoundError() *NotFound {
	notFoundErrorOnce.Do(func() {
		notFoundErrorInstance = &NotFound{
			message: "Internal error",
		}
	})
	return notFoundErrorInstance
}

func (err *NotFound) Error() string {
	return err.message
}
