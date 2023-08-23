package handler

import (
	"errors"
	"log/slog"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	ConnectEvent = shadiaosocketio.OnConnection
)

type ConnectState interface {
	Connected() (bool, error)
	SetConnected(bool) error
}

// Connect is a handler for the connect event.
type Connect struct {
	state ConnectState
}

// NewConnect returns a new Connect handler.
func NewConnect(state ConnectState) *Connect {
	return &Connect{state: state}
}

// Event returns the event name.
func (c *Connect) Event() string {
	return ConnectEvent
}

// Register registers the handler.
func (c *Connect) Register(h HandlerRegistrator) error {
	return h.On(ConnectEvent, c.Callback)
}

// Occurred returns true if the event has occurred at least once.
func (c *Connect) Occurred() bool {
	_, err := c.state.Connected()
	return !errors.Is(err, state.ErrNotSetYet)
}

// Callback handles the event.
func (c *Connect) Callback(ch *shadiaosocketio.Channel) error {
	slog.Debug("received connect event")

	if ch == nil {
		return errors.New("nil channel")
	}

	if err := c.state.SetConnected(true); err != nil {
		return err
	}

	return nil
}
