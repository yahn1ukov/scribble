package services

import "errors"

var (
	ErrTitleIsRequired  = errors.New("title is required")
	ErrNoFieldsToUpdate = errors.New("no fields to update")
)
