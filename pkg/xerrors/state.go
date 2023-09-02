package xerrors

import (
	"errors"
	"fmt"
)

var (
	// xerrors.ErrStateNil is returned when the state is nil, e.g. when the client is not
	// properly initialized.
	ErrStateNil = errors.New("state is nil")

	// xerrors.ErrNotSetYet is returned when the value is not set yet - this may be the case
	// if the client has not yet received the event that sets the value.
	ErrNotSetYet = errors.New("value not set yet")
)

// ErrNotFound is returned when a resource for a given ID is not found in the current state cache.
type ErrNotFound struct {
	Kind string
	Id   int
}

// NewErrNotFound returns a new ErrNotFound.
func NewErrNotFound(kind string, id int) *ErrNotFound {
	return &ErrNotFound{
		Kind: kind,
		Id:   id,
	}
}

// Error returns the error message.
func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s with id %d not found", e.Kind, e.Id)
}
