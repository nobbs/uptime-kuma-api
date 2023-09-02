package handler

import (
	"fmt"
	"log/slog"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/nobbs/uptime-kuma-api/pkg/xerrors"
)

const (
	MonitorListEvent = "monitorList"
)

type MonitorState interface {
	SetMonitors(monitors map[int]*state.Monitor) (err error)
	HasSeen(event string) (seen bool)
	MarkSeen(event string)
}

type MonitorList struct {
	state MonitorState
}

func NewMonitorList(state MonitorState) *MonitorList {
	return &MonitorList{state: state}
}

func (ml MonitorList) Event() string {
	return MonitorListEvent
}

func (ml MonitorList) Register(h HandlerRegistrator) error {
	return h.On(MonitorListEvent, ml.Callback)
}

func (ml MonitorList) Occurred() bool {
	return ml.state.HasSeen(MonitorListEvent)
}

func (ml MonitorList) Callback(ch *shadiaosocketio.Channel, data any) error {
	slog.Info("MonitorList callback")
	ml.state.MarkSeen(MonitorListEvent)

	// assert data type
	typedData, ok := data.(map[string]any)
	if !ok {
		return xerrors.NewErrInvalidDataType("map[string]any", data)
	}

	// decode data into struct
	monitors := make(map[int]*state.Monitor)
	if err := utils.DecodeMap(typedData, monitors); err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	// set monitors
	if err := ml.state.SetMonitors(monitors); err != nil {
		return err
	}

	return nil
}
