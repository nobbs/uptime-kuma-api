package handler

import (
	"log/slog"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	ErrorEvent = shadiaosocketio.OnError
)

type Error struct {
	state *state.State
}

func NewError(state *state.State) *Error {
	return &Error{state: state}
}

func (e *Error) Event() string {
	return ErrorEvent
}

func (e *Error) Register(h HandlerRegistrator) error {
	return h.On(ErrorEvent, e.Callback)
}

func (e *Error) Occurred() bool {
	// TODO: implement some kind of error handling here
	return false
}

func (e *Error) Callback(ch *shadiaosocketio.Channel, data any) error {
	slog.Warn("received error event", slog.Any("data", data))
	return nil
}
