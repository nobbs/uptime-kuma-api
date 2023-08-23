package handler

import (
	"log/slog"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/gorilla/websocket"
)

const (
	DisconnectEvent = shadiaosocketio.OnDisconnection
)

type DisconnectState interface {
	SetConnected(bool) error
}

// Disconnect is a handler for the disconnect event.
type Disconnect struct {
	state DisconnectState
}

// NewDisconnect returns a new Disconnect handler.
func NewDisconnect(state DisconnectState) *Disconnect {
	return &Disconnect{state: state}
}

// Event returns the event name.
func (d *Disconnect) Event() string {
	return DisconnectEvent
}

// Register registers the handler.
func (d *Disconnect) Register(h HandlerRegistrator) error {
	return h.On(DisconnectEvent, d.Callback)
}

// Occurred returns true if the event has occurred at least once.
func (d *Disconnect) Occurred() bool {
	// TODO: implement either reconnect or exit logic
	return false
}

// Callback handles the event.
func (d *Disconnect) Callback(ch *shadiaosocketio.Channel, reason websocket.CloseError) error {
	slog.Debug("received disconnect event", slog.Int("code", reason.Code), slog.String("text", reason.Text))

	if ch == nil {
		return nil
	}

	if err := d.state.SetConnected(false); err != nil {
		return err
	}

	return nil
}
