package ierr

import (
	"errors"
	"fmt"
)

func NewError(format string, arg ...interface{}) error {
	return errors.New(fmt.Sprintf(format, arg...))
}

func IErr(code int) error {
	return NewError(ErrorMessage[code])
}
