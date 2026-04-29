package domain

import (
	"errors"
)

type AppError struct {
	err        error
	statusCode int
}

func NewAppError(errCode int, errs ...error) *AppError {
	return &AppError{
		err:        errors.Join(errs...),
		statusCode: errCode,
	}
}

func (a *AppError) Error() string {
	if a.err != nil {
		return a.err.Error()
	}

	return ""
}

func (a *AppError) Unwrap() error {
	return a.err
}
