package state

import (
	"errors"
	"fmt"
)

var (
	// ErrStateNil is returned when the state is nil, e.g. when the client is not
	// properly initialized.
	ErrStateNil = errors.New("state is nil")

	// ErrNotSetYet is returned when the value is not set yet - this may be the case
	// if the client has not yet received the event that sets the value.
	ErrNotSetYet = errors.New("value not set yet")
)

// ErrNotFound is returned when a resource for a given ID is not found.
type ErrNotFound struct {
	Kind string
	Id   int
}

func NewErrNotFound(kind string, id int) *ErrNotFound {
	return &ErrNotFound{
		Kind: kind,
		Id:   id,
	}
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s with id %d not found", e.Kind, e.Id)
}
