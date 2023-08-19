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

// Connect is a handler for the connect event.
type Connect struct {
	state *state.State
}

// NewConnect returns a new Connect handler.
func NewConnect(state *state.State) *Connect {
	return &Connect{state: state}
}

// Event returns the event name.
func (c *Connect) Event() string {
	return ConnectEvent
}

// Register registers the handler.
func (c *Connect) Register(h HandlerRegistrator) error {
	fn := func(ch *shadiaosocketio.Channel) error {
		slog.Debug("received connect event")

		if ch == nil {
			return errors.New("nil channel")
		}

		c.state.SetConnected(true)
		return nil
	}

	return h.On(ConnectEvent, fn)
}

// Occured returns true if the event has occured at least once.
func (c *Connect) Occured() bool {
	_, err := c.state.Connected()
	return !errors.Is(err, state.ErrNotSetYet)
}
