package xerrors

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrTimeout = errors.New("timeout occurred")

type ErrInvalidDataType struct {
	Expected string
	Actual   string
}

func NewErrInvalidDataType(expected string, actual any) ErrInvalidDataType {
	return ErrInvalidDataType{
		Actual: reflect.TypeOf(actual).String(),
	}
}

func (e ErrInvalidDataType) Error() string {
	return fmt.Sprintf("invalid data type: expected %s, got %s", e.Expected, e.Actual)
}
