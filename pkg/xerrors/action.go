package xerrors

import (
	"errors"
	"fmt"
)

// ErrInvalidResponse is returned when the response from the server is invalid.
var ErrInvalidResponse = errors.New("invalid response")

// Err2faTokenRequired is returned when a 2fa token is required to login.
var Err2faTokenRequired = errors.New("2fa token required")

type ErrAwaitFailed struct {
	Event string
	err   error
}

type ErrActionFailed struct {
	Action string
	Msg    string
}

type ErrLoginFailed struct {
	Msg string
}

func NewErrActionFailed(action, msg string) ErrActionFailed {
	return ErrActionFailed{Action: action, Msg: msg}
}

func (e ErrActionFailed) Error() string {
	return fmt.Sprintf("action failed: %s", e.Msg)
}

func NewErrLoginFailed(msg string) ErrLoginFailed {
	return ErrLoginFailed{Msg: msg}
}

func (e ErrLoginFailed) Error() string {
	return fmt.Sprintf("auth failed: %s", e.Msg)
}

func NewErrAwaitFailed(event string, err error) ErrAwaitFailed {
	return ErrAwaitFailed{Event: event, err: err}
}

func (e ErrAwaitFailed) Error() string {
	return fmt.Sprintf("await failed for event %s: %s", e.Event, e.err)
}
