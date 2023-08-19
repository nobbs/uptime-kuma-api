package handler

import (
	"log/slog"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	MessageEvent = shadiaosocketio.OnMessage
)

type Message struct {
	state *state.State
}

func NewMessage(state *state.State) *Message {
	return &Message{state}
}

func (m *Message) Event() string {
	return MessageEvent
}

func (m *Message) Register(h HandlerRegistrator) error {
	return h.On(MessageEvent, m.Callback)
}

func (m *Message) Occured() bool {
	// TODO: implement some kind of message handling here
	return false
}

func (m *Message) Callback(ch *shadiaosocketio.Channel, data any) error {
	slog.Info("received message event", slog.Any("data", data))
	return nil
}
