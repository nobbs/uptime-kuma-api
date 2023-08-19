package handler

import (
	"errors"
	"log/slog"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	AutoLoginEvent = "autoLogin"
)

type AutoLogin struct {
	state *state.State
}

func NewAutoLogin(state *state.State) *AutoLogin {
	return &AutoLogin{state: state}
}

func (al *AutoLogin) Event() string {
	return AutoLoginEvent
}

func (al *AutoLogin) Register(h HandlerRegistrator) error {
	fn := func(ch *shadiaosocketio.Channel) error {
		slog.Debug("received auto login event")

		if err := al.state.SetAutoLogin(true); err != nil {
			return err
		}

		return nil
	}

	return h.On(AutoLoginEvent, fn)
}

func (al *AutoLogin) Occured() bool {
	_, err := al.state.AutoLogin()
	return !errors.Is(err, state.ErrNotSetYet)
}
