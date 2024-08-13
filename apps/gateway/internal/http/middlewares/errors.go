package middlewares

import "errors"

var (
	ErrInvalidFormat = errors.New("invalid format")
	ErrAccessDenied  = errors.New("access denied")
	ErrUnauthorized  = errors.New("unauthorized")
)
