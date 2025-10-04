package go_flags

import (
	"errors"
)

const (
	ErrInvalidValue = "invalid value %q, allowed values are: %s"
)

var (
	ErrNilFlag = errors.New("flag cannot be nil")
)
