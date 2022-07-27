package service_errors

import (
	"errors"
	"log"
	custom_errors "relembrando_2/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(err error) error {
	if err == nil {
		return nil
	} else if errors.Is(err, custom_errors.NewNotFoundError()) {
		return handleNotFoundError(err)
	} else if errors.Is(err, custom_errors.NewInvalidIdError()) {
		return handleInvalidIdError(err)
	} else {
		return handleUnknownError(err)
	}
}

func handleUnknownError(err error) error {
	log.Print(err.Error())
	return status.Error(codes.Unknown, "Unknown error")
}

func handleNotFoundError(err error) error {
	log.Print(err.Error())
	return status.Error(codes.NotFound, "Not found")
}

func handleInvalidIdError(err error) error {
	log.Print(err.Error())
	return status.Error(codes.InvalidArgument, "Invalid id")
}
