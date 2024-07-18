package grpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errParse = errors.New("grpc parsing error")
)

type Error struct {
	code codes.Code
	err  error
}

func (e *Error) Code() codes.Code {
	return e.code
}

func (e *Error) Error() error {
	return e.err
}

func CreateError(code codes.Code, err error) *Error {
	return &Error{
		code: code,
		err:  err,
	}
}

func ParseError(err error) *Error {
	status, ok := status.FromError(err)
	if !ok {
		return &Error{
			code: codes.Internal,
			err:  errParse,
		}
	}

	return &Error{
		code: status.Code(),
		err:  fmt.Errorf("%s", status.Message()),
	}
}
