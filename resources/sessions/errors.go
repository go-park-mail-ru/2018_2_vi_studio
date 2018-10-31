package sessions

import "errors"

var (
	ErrUnknown = errors.New("unknown error")
	ErrNotFound = errors.New("not found error")
)