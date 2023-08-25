package action

import (
	"fmt"

	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	getMonitorAction     = "getMonitor"
	getMonitorListAction = "getMonitorList"

	addMonitorAction    = "add"
	editMonitorAction   = "editMonitor"
	deleteMonitorAction = "deleteMonitor"
	pauseMonitorAction  = "pauseMonitor"
	resumeMonitorAction = "resumeMonitor"

	getMonitorBeatsAction = "getMonitorBeats"
	clearEventsAction     = "clearEvents"
	clearHeartbeatsAction = "clearHeartbeats"
	clearStatisticsAction = "clearStatistics"

	addMonitorTagAction    = "addMonitorTag"
	editMonitorTagAction   = "editMonitorTag"
	deleteMonitorTagAction = "deleteMonitorTag"
)

type getMonitorListResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type getMonitorResponse struct {
	Ok      bool           `mapstructure:"ok"`
	Msg     *string        `mapstructure:"msg"`
	Monitor *state.Monitor `mapstructure:"monitor"`
}

type addMonitorResponse struct {
	Ok        bool    `mapstructure:"ok"`
	Msg       *string `mapstructure:"msg"`
	MonitorId *int    `mapstructure:"monitorID"`
}

type editMonitorResponse struct {
	Ok        bool    `mapstructure:"ok"`
	Msg       *string `mapstructure:"msg"`
	MonitorId *int    `mapstructure:"monitorID"`
}

type pauseMonitorResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type resumeMonitorResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type deleteMonitorResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type getMonitorBeatsResponse struct {
	Ok   bool              `mapstructure:"ok"`
	Msg  *string           `mapstructure:"msg"`
	Data []state.Heartbeat `mapstructure:"data"`
}

type clearEventsResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type clearHeartbeatsResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type clearStatisticsResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type addMonitorTagResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type editMonitorTagResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

type deleteMonitorTagResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

// GetMonitorList triggers the server to emit the monitor list event. The monitor list event
// contains a list of all monitors in the Uptime Kuma instance and is handled by the monitorList
// event handler.
func GetMonitorList(c StatefulEmiter) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(getMonitorListAction, defaultEmitTimeout)
	if err != nil {
		return NewErrActionFailed(getMonitorListAction, err.Error())
	}

	// unmarshal raw response data
	data := &getMonitorListResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(getMonitorListAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(getMonitorListAction, *data.Msg)
	}

	return nil
}

// GetMonitor requests the data of a specific monitor from the Uptime Kuma instance that is
// send as a response. The monitor data is also stored or updated in the client state.
func GetMonitor(c StatefulEmiter, monitorId int) (*state.Monitor, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return nil, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(getMonitorAction, defaultEmitTimeout, monitorId)
	if err != nil {
		return nil, NewErrActionFailed(getMonitorAction, err.Error())
	}

	// unmarshal raw response data
	data := &getMonitorResponse{}
	if err := decode(response, data); err != nil {
		return nil, NewErrActionFailed(getMonitorAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return nil, NewErrActionFailed(getMonitorAction, *data.Msg)
	}

	// update state
	if err := c.State().SetMonitor(monitorId, data.Monitor); err != nil {
		return nil, err
	}

	return data.Monitor, nil
}

// AddMonitor adds a new monitor to the Uptime Kuma instance.
func AddMonitor(c StatefulEmiter, monitor *state.Monitor) (int, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return 0, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(addMonitorAction, defaultEmitTimeout, monitor)
	if err != nil {
		return 0, NewErrActionFailed(addMonitorAction, err.Error())
	}

	// unmarshal raw response data
	data := &addMonitorResponse{}
	if err := decode(response, data); err != nil {
		return 0, NewErrActionFailed(addMonitorAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return 0, NewErrActionFailed(addMonitorAction, *data.Msg)
	}

	return *data.MonitorId, nil
}

// EditMonitor edits an existing monitor in the Uptime Kuma instance.
func EditMonitor(c StatefulEmiter, monitor *state.Monitor) (int, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return 0, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(editMonitorAction, defaultEmitTimeout, monitor)
	if err != nil {
		return 0, NewErrActionFailed(editMonitorAction, err.Error())
	}

	// unmarshal raw response data
	data := &editMonitorResponse{}
	if err := decode(response, data); err != nil {
		return 0, NewErrActionFailed(editMonitorAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return 0, NewErrActionFailed(editMonitorAction, *data.Msg)
	}

	return *data.MonitorId, nil
}

// PauseMonitor pauses a monitor in the Uptime Kuma instance.
func PauseMonitor(c StatefulEmiter, monitorId int) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(pauseMonitorAction, defaultEmitTimeout, monitorId)
	if err != nil {
		return NewErrActionFailed(pauseMonitorAction, err.Error())
	}

	// unmarshal raw response data
	data := &pauseMonitorResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(pauseMonitorAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(pauseMonitorAction, *data.Msg)
	}

	return nil
}

// ResumeMonitor resumes a monitor in the Uptime Kuma instance.
func ResumeMonitor(c StatefulEmiter, monitorId int) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(resumeMonitorAction, defaultEmitTimeout, monitorId)
	if err != nil {
		return NewErrActionFailed(resumeMonitorAction, err.Error())
	}

	// unmarshal raw response data
	data := &resumeMonitorResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(resumeMonitorAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(resumeMonitorAction, *data.Msg)
	}

	return nil
}

// DeleteMonitor deletes a monitor in the Uptime Kuma instance.
func DeleteMonitor(c StatefulEmiter, monitorId int) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(deleteMonitorAction, defaultEmitTimeout, monitorId)
	if err != nil {
		return fmt.Errorf("%s: %w", deleteMonitorAction, err)
	}

	// unmarshal raw response data
	data := &deleteMonitorResponse{}
	if err := decode(response, data); err != nil {
		return fmt.Errorf("%s: %w", deleteMonitorAction, err)
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(deleteMonitorAction, *data.Msg)
	}

	return nil
}

// GetMonitorBeats requests the heartbeats of a specific monitor and period of hours from the Uptime
// Kuma instance that are send as a response.
func GetMonitorBeats(c StatefulEmiter, monitorId int, hours int) ([]state.Heartbeat, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return nil, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(getMonitorBeatsAction, defaultEmitTimeout, monitorId, hours)
	if err != nil {
		return nil, NewErrActionFailed(getMonitorBeatsAction, err.Error())
	}

	// unmarshal raw response data
	data := &getMonitorBeatsResponse{}
	if err := decode(response, data); err != nil {
		return nil, NewErrActionFailed(getMonitorBeatsAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return nil, NewErrActionFailed(getMonitorBeatsAction, *data.Msg)
	}

	return data.Data, nil
}

// ClearEvents clears the events of a monitor in the Uptime Kuma instance.
func ClearEvents(c StatefulEmiter, monitorId int) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(clearEventsAction, defaultEmitTimeout, monitorId)
	if err != nil {
		return NewErrActionFailed(clearEventsAction, err.Error())
	}

	// unmarshal raw response data
	data := &clearEventsResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(clearEventsAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(clearEventsAction, *data.Msg)
	}

	return nil
}

// ClearHeartbeats clears the heartbeats of a monitor in the Uptime Kuma instance.
func ClearHeartbeats(c StatefulEmiter, monitorId int) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(clearHeartbeatsAction, defaultEmitTimeout, monitorId)
	if err != nil {
		return NewErrActionFailed(clearHeartbeatsAction, err.Error())
	}

	// unmarshal raw response data
	data := &clearHeartbeatsResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(clearHeartbeatsAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(clearHeartbeatsAction, *data.Msg)
	}

	return nil
}

// ClearStatistics clears the events and heartbeats of all monitors in the Uptime Kuma instance.
func ClearStatistics(c StatefulEmiter) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(clearStatisticsAction, defaultEmitTimeout)
	if err != nil {
		return NewErrActionFailed(clearStatisticsAction, err.Error())
	}

	// unmarshal raw response data
	data := &clearStatisticsResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(clearStatisticsAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(clearStatisticsAction, *data.Msg)
	}

	return nil
}

// AddMonitorTag adds a tag to a monitor in the Uptime Kuma instance with the given value.
func AddMonitorTag(c StatefulEmiter, monitorId, tagId int, value string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return fmt.Errorf("%s: %w", handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(addMonitorTagAction, defaultEmitTimeout, tagId, monitorId, value)
	if err != nil {
		return fmt.Errorf("%s: %w", addMonitorTagAction, err)
	}

	// unmarshal raw response data
	data := &addMonitorTagResponse{}
	if err := decode(response, data); err != nil {
		return fmt.Errorf("%s: %w", addMonitorTagAction, err)
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(addMonitorTagAction, *data.Msg)
	}

	return nil
}

// EditMonitorTag edits a tag of a monitor in the Uptime Kuma instance with the given value.
func EditMonitorTag(c StatefulEmiter, monitorId, tagId int, value string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return fmt.Errorf("%s: %w", handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(editMonitorTagAction, defaultEmitTimeout, tagId, monitorId, value)
	if err != nil {
		return fmt.Errorf("%s: %w", editMonitorTagAction, err)
	}

	// unmarshal raw response data
	data := &editMonitorTagResponse{}
	if err := decode(response, data); err != nil {
		return fmt.Errorf("%s: %w", editMonitorTagAction, err)
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(editMonitorTagAction, *data.Msg)
	}

	return nil
}

// DeleteMonitorTag deletes a tag of a monitor in the Uptime Kuma instance.
func DeleteMonitorTag(c StatefulEmiter, monitorId, tagId int, value string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return fmt.Errorf("%s: %w", handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(deleteMonitorTagAction, defaultEmitTimeout, tagId, monitorId, value)
	if err != nil {
		return fmt.Errorf("%s: %w", deleteMonitorTagAction, err)
	}

	// unmarshal raw response data
	data := &deleteMonitorTagResponse{}
	if err := decode(response, data); err != nil {
		return fmt.Errorf("%s: %w", deleteMonitorTagAction, err)
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(deleteMonitorTagAction, *data.Msg)
	}

	return nil
}
