package action

import (
	"fmt"

	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	// getMonitorAction     = "getMonitor"
	// getMonitorListAction = "getMonitorList"

	pauseMonitorAction  = "pauseMonitor"
	resumeMonitorAction = "resumeMonitor"

	addMonitorAction    = "add"
	editMonitorAction   = "editMonitor"
	deleteMonitorAction = "deleteMonitor"

	clearEventsAction     = "clearEvents"
	clearHeartbeatsAction = "clearHeartbeats"
	clearStatisticsAction = "clearStatistics"
)

// type getMonitorResponse struct {
// 	Ok      bool           `mapstructure:"ok"`
// 	Msg     *string        `mapstructure:"msg"`
// 	Monitor *state.Monitor `mapstructure:"monitor"`
// }

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

// func GetMonitorList(c StatefulEmiter) error {
// 	// call action
// 	response, err := call(c, getMonitorListAction)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", getMonitorListAction, err)
// 	}

// 	// unmarshal raw response data
// 	data := &monitorGenericResponse{}
// 	err = decode(response, data)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", getMonitorListAction, err)
// 	}

// 	// check if action was successful
// 	if !data.Ok {
// 		return NewErrActionFailed(getMonitorListAction, *data.Msg)
// 	}

// 	return nil
// }

// func GetMonitor(c StatefulEmiter, monitorId int) (*state.Monitor, error) {
// 	// call action
// 	response, err := call(c, getMonitorAction, monitorId)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", getMonitorAction, err)
// 	}

// 	// unmarshal raw response data
// 	data := &getMonitorResponse{}
// 	err = decode(response, data)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", getMonitorAction, err)
// 	}

// 	// check if action was successful
// 	if !data.Ok {
// 		return nil, NewErrActionFailed(getMonitorAction, *data.Msg)
// 	}

// 	return data.Monitor, nil
// }

// AddMonitor adds a new monitor to the Uptime Kuma instance.
func AddMonitor(c StatefulEmiter, monitor *state.Monitor) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(addMonitorAction, defaultEmitTimeout, monitor)
	if err != nil {
		return NewErrActionFailed(addMonitorAction, err.Error())
	}

	// unmarshal raw response data
	data := &addMonitorResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(addMonitorAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(addMonitorAction, *data.Msg)
	}

	return nil
}

// EditMonitor edits an existing monitor in the Uptime Kuma instance.
func EditMonitor(c StatefulEmiter, monitor *state.Monitor) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(editMonitorAction, defaultEmitTimeout, monitor)
	if err != nil {
		return NewErrActionFailed(editMonitorAction, err.Error())
	}

	// unmarshal raw response data
	data := &editMonitorResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(editMonitorAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(editMonitorAction, *data.Msg)
	}

	return nil
}

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
	if err = decode(response, data); err != nil {
		return NewErrActionFailed(clearStatisticsAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(clearStatisticsAction, *data.Msg)
	}

	return nil
}

// func ClearEvents(c StatefulEmiter, monitorId int) error {
// 	// call action
// 	response, err := call(c, clearEventsAction, monitorId)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", clearEventsAction, err)
// 	}

// 	// unmarshal raw response data
// 	data := &monitorGenericResponse{}
// 	err = decode(response, data)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", clearEventsAction, err)
// 	}

// 	// check if action was successful
// 	if !data.Ok {
// 		return NewErrActionFailed(clearEventsAction, *data.Msg)
// 	}

// 	return nil
// }
