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

type AutoLoginState interface {
	AutoLogin() (bool, error)
	SetAutoLogin(bool) error
}

type AutoLogin struct {
	state AutoLoginState
}

func NewAutoLogin(state AutoLoginState) *AutoLogin {
	return &AutoLogin{state: state}
}

func (al *AutoLogin) Event() string {
	return AutoLoginEvent
}

func (al *AutoLogin) Register(h HandlerRegistrator) error {
	return h.On(AutoLoginEvent, al.Callback)
}

func (al *AutoLogin) Occured() bool {
	_, err := al.state.AutoLogin()
	return !errors.Is(err, state.ErrNotSetYet)
}

func (al *AutoLogin) Callback(ch *shadiaosocketio.Channel) error {
	slog.Debug("received auto login event")

	if err := al.state.SetAutoLogin(true); err != nil {
		return err
	}

	return nil
}
