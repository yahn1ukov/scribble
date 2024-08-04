package repositories

import "errors"

var (
	ErrNotFound      = errors.New("notebook not found")
	ErrAlreadyExists = errors.New("notebook already exists")
)
