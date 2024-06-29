package domain

import "errors"

var (
	ErrNotFound      = errors.New("notebook not found")
	ErrAlreadyExists = errors.New("notebook already exists")
	ErrTitleRequired = errors.New("notebook title is required")
)
