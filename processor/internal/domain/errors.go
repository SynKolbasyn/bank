package domain

import (
	"errors"
)

type AppError struct {
	err     error
	errCode int
}

func NewAppError(errCode int, errs ...error) *AppError {
	return &AppError{
		err:     errors.Join(errs...),
		errCode: errCode,
	}
}

func (a *AppError) Error() string {
	return a.err.Error()
}
