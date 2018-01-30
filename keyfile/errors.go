package keyfile

import (
	"errors"
)

var (
	ErrBadEscapeSequence     = errors.New("Bad Escape Sequence")
	ErrUnexpectedEndOfString = errors.New("Unexpected end of string")
)
