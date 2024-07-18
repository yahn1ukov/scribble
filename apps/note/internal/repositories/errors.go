package repositories

import "errors"

var (
	ErrNotFound      = errors.New("note not found")
	ErrTitleRequired = errors.New("note title is required")
)
