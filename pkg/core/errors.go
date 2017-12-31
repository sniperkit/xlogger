package core

import (
	"errors"
)

var (
	ErrIllegalBackend  = errors.New("illegal log backend")
	ErrIllegalLevel    = errors.New("illegal log level")
	ErrIllegalEncoding = errors.New("illegal log encoding")
)
