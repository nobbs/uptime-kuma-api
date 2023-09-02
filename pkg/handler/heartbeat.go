package handler

import (
	"errors"
	"fmt"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/nobbs/uptime-kuma-api/pkg/xerrors"
)

const (
	HeartbeatEvent              = "heartbeat"
	HeartbeatListEvent          = "heartbeatList"
	ImportantHeartbeatListEvent = "importantHeartbeatList"
)

type HeartbeatState interface {
	Heartbeats(monitorId int) (beats []state.Heartbeat, err error)
	AppendHeartbeat(beat *state.Heartbeat) (err error)
}

type HeartbeatListState interface {
	Heartbeats(monitorId int) (beats []state.Heartbeat, err error)
	SetHeartbeats(monitorId int, beats []state.Heartbeat, overwrite bool) (err error)
}

type ImportantHeartbeatListState interface {
	ImportantHeartbeats(monitorId int) (beats []state.Heartbeat, err error)
	SetImportantHeartbeats(monitorId int, beats []state.Heartbeat, overwrite bool) (err error)
}

type Heartbeat struct {
	state HeartbeatState
}

type HeartbeatList struct {
	state HeartbeatListState
}

type ImportantHeartbeatList struct {
	state ImportantHeartbeatListState
}

func NewHeartbeat(state HeartbeatState) *Heartbeat {
	return &Heartbeat{state: state}
}

func NewHeartbeatList(state HeartbeatListState) *HeartbeatList {
	return &HeartbeatList{state: state}
}

func NewImportantHeartbeatList(state ImportantHeartbeatListState) *ImportantHeartbeatList {
	return &ImportantHeartbeatList{state: state}
}

func (hn *Heartbeat) Event() string {
	return HeartbeatEvent
}

func (hn *HeartbeatList) Event() string {
	return HeartbeatListEvent
}

func (hn *ImportantHeartbeatList) Event() string {
	return ImportantHeartbeatListEvent
}

func (hn *Heartbeat) Register(h HandlerRegistrator) error {
	return h.On(HeartbeatEvent, hn.Callback)
}

func (hn *HeartbeatList) Register(h HandlerRegistrator) error {
	return h.On(HeartbeatListEvent, hn.Callback)
}

func (hn *ImportantHeartbeatList) Register(h HandlerRegistrator) error {
	return h.On(ImportantHeartbeatListEvent, hn.Callback)
}

func (hn *Heartbeat) Occurred() bool {
	_, err := hn.state.Heartbeats(0)
	return err == nil || !errors.Is(err, xerrors.ErrNotSetYet)
}

func (hn *HeartbeatList) Occurred() bool {
	_, err := hn.state.Heartbeats(0)
	return err == nil || !errors.Is(err, xerrors.ErrNotSetYet)
}

func (hn *ImportantHeartbeatList) Occurred() bool {
	_, err := hn.state.ImportantHeartbeats(0)
	return err == nil || !errors.Is(err, xerrors.ErrNotSetYet)
}

func (hn *Heartbeat) Callback(h *shadiaosocketio.Channel, data any) error {
	// assert data type
	typedData, ok := data.(map[string]any)
	if !ok {
		return xerrors.NewErrInvalidDataType("map[string]any", data)
	}

	// decode data into struct
	heartbeat := &state.Heartbeat{}
	if err := utils.Decode(typedData, heartbeat); err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	// store heartbeat
	if err := hn.state.AppendHeartbeat(heartbeat); err != nil {
		return err
	}

	return nil
}

func (hn *HeartbeatList) Callback(h *shadiaosocketio.Channel, id any, result []any, overwrite any) error {
	// parse monitorId and overwrite flag first
	data := map[string]any{
		"monitorId": id,
		"overwrite": overwrite,
	}

	// decode monitorId and overwrite into data
	response := &struct {
		MonitorId int  `mapstructure:"monitorId"`
		Overwrite bool `mapstructure:"overwrite"`
	}{}
	if err := utils.Decode(data, response); err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	// decode heartbeats into slice of heartbeats
	heartbeats := make([]state.Heartbeat, 0, len(result))

	heartbeats, err := utils.DecodeSlice(result, heartbeats)
	if err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	// set heartbeats
	if err := hn.state.SetHeartbeats(response.MonitorId, heartbeats, response.Overwrite); err != nil {
		return err
	}

	return nil
}

func (hn *ImportantHeartbeatList) Callback(h *shadiaosocketio.Channel, id any, result []any, overwrite any) error {
	// parse monitorId and overwrite flag first
	data := map[string]any{
		"monitorId": id,
		"overwrite": overwrite,
	}

	// decode monitorId and overwrite into data
	response := &struct {
		MonitorId int  `mapstructure:"monitorId"`
		Overwrite bool `mapstructure:"overwrite"`
	}{}
	if err := utils.Decode(data, response); err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	// decode heartbeats into slice of heartbeats
	heartbeats := make([]state.Heartbeat, 0, len(result))

	heartbeats, err := utils.DecodeSlice(result, heartbeats)
	if err != nil {
		return fmt.Errorf("decode failed: %w", err)
	}

	// set heartbeats
	if err := hn.state.SetImportantHeartbeats(response.MonitorId, heartbeats, response.Overwrite); err != nil {
		return err
	}

	return nil
}
