package http

import "errors"

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrInvalidFormat = errors.New("invalid format")
)
