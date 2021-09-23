package apierr

import "errors"

type ApiError struct {
	code int
	error
}

func NewApiError(description string, code int) ApiError {
	err := ApiError{
		code: code,
	}

	err.error = errors.New(description)

	return err
}
