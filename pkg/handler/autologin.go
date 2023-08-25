package handler

import (
	"log/slog"

	"github.com/Baiguoshuai1/shadiaosocketio"
)

const (
	AutoLoginEvent = "autoLogin"
)

type AutoLoginState interface {
	SetAutoLogin(autoLogin bool) (err error)
	HasSeen(event string) (seen bool)
	MarkSeen(event string)
}

type AutoLogin struct {
	state AutoLoginState
}

// NewAutoLogin creates a new AutoLogin handler.
func NewAutoLogin(state AutoLoginState) *AutoLogin {
	return &AutoLogin{state: state}
}

// Event returns the event name.
func (al *AutoLogin) Event() string {
	return AutoLoginEvent
}

// Register registers the handler with the client.
func (al *AutoLogin) Register(h HandlerRegistrator) error {
	return h.On(AutoLoginEvent, al.Callback)
}

// Occurred returns true if the event has occurred, false otherwise. Required in some places to
// make sure a specific event has been sent before continuing.
func (al *AutoLogin) Occurred() bool {
	return al.state.HasSeen(AutoLoginEvent)
}

// Callback is the function that is called when the event is received.
func (al *AutoLogin) Callback(ch *shadiaosocketio.Channel) error {
	slog.Info("AutoLogin callback")
	al.state.MarkSeen(AutoLoginEvent)

	if err := al.state.SetAutoLogin(true); err != nil {
		return err
	}

	return nil
}
