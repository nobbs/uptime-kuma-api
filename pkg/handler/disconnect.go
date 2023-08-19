package handler

import (
	"log/slog"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/gorilla/websocket"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	DisconnectEvent = shadiaosocketio.OnDisconnection
)

type Disconnect struct {
	state *state.State
}

func NewDisconnect(state *state.State) *Disconnect {
	return &Disconnect{state: state}
}

func (d *Disconnect) Event() string {
	return DisconnectEvent
}

func (d *Disconnect) Register(h HandlerRegistrator) error {
	fn := func(ch *shadiaosocketio.Channel, reason websocket.CloseError) error {
		slog.Debug("received disconnect event", slog.Int("code", reason.Code), slog.String("text", reason.Text))

		if ch == nil {
			return nil
		}

		if err := d.state.SetConnected(false); err != nil {
			return err
		}

		return nil
	}

	return h.On(DisconnectEvent, fn)
}

func (d *Disconnect) Occured() bool {
	// TODO: implement either reconnect or exit logic
	return false
}
