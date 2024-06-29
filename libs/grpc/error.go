package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	code    codes.Code
	message string
}

func (e *Error) Code() codes.Code {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func ParseError(err error) *Error {
	status, ok := status.FromError(err)
	if !ok {
		return &Error{
			code:    codes.Internal,
			message: "grpc parsing error",
		}
	}

	return &Error{
		code:    status.Code(),
		message: status.Message(),
	}
}
