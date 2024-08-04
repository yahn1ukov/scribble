package services

import "errors"

var (
	ErrEmailIsRequired       = errors.New("email is required")
	ErrPasswordIsRequired    = errors.New("password is required")
	ErrInvalidPassword       = errors.New("password is invalid")
	ErrOldPasswordIsRequired = errors.New("old password is required")
	ErrNewPasswordIsRequired = errors.New("new password is required")
	ErrPasswordsAreSame      = errors.New("passwords can't be the same")
	ErrNoFieldsToUpdate      = errors.New("no fields to update")
)
