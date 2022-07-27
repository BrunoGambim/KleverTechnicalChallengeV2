package mongodb_errors

import (
	"errors"
	custom_errors "relembrando_2/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handle(err error) error {
	if err == nil {
		return nil
	} else if custom_errors.IsACustomError(err) {
		return err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return custom_errors.NewNotFoundError()
	} else if errors.Is(err, primitive.ErrInvalidHex) {
		return custom_errors.NewInvalidIdError()
	} else {
		return custom_errors.NewNotFoundError()
	}
}
