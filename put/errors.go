package put

import (
	"errors"
)

var (
	ErrZeroArgument     = errors.New("zero argument")
	ErrNegativeArgument = errors.New("negative argument")
)
