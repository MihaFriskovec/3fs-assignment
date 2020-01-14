package errors

import errors "github.com/MihaFriskovec/3fs-assignment/errors/models"

func NewError(errorType string, message string, status int) errors.Error {
	err := errors.Error{}

	err.Type = errorType
	err.Message = message
	err.Status = status

	return err
}
