package handler

import (
	"fmt"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

const (
	InfoEvent = "info"
)

type InfoState interface {
	Info() (*state.Info, error)
	SetInfo(*state.Info) error
}

type Info struct {
	state InfoState
}

func NewInfo(state InfoState) *Info {
	return &Info{state: state}
}

func (i *Info) Event() string {
	return InfoEvent
}

func (i *Info) Register(h HandlerRegistrator) error {
	return h.On(InfoEvent, i.Callback)
}

func (i *Info) Occured() bool {
	_, err := i.state.Info()
	return err == nil
}

func (i *Info) Callback(ch *shadiaosocketio.Channel, data any) error {
	// assert data type
	typedData, ok := data.(map[string]any)
	if !ok {
		return NewErrInvalidDataType("map[string]any", data)
	}

	// decode data into struct
	info := &state.Info{}
	if err := utils.Decode(typedData, info); err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	// set info
	if err := i.state.SetInfo(info); err != nil {
		return err
	}

	return nil
}
