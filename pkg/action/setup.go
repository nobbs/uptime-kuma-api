package action

import (
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

const (
	needSetupAction = "needSetup"
	setupAction     = "setup"
)

type setupResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

func NeedSetup(c StatefulEmiter) (bool, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return false, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(needSetupAction, defaultEmitTimeout)
	if err != nil {
		return false, NewErrActionFailed(needSetupAction, err.Error())
	}

	// unmarshal raw response data
	data := utils.NewBool(false)
	if err := decode(response, data); err != nil {
		return false, NewErrActionFailed(needSetupAction, err.Error())
	}

	return *data, nil
}

func Setup(c StatefulEmiter, username, password string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(setupAction, defaultEmitTimeout, username, password)
	if err != nil {
		return NewErrActionFailed(setupAction, err.Error())
	}

	// unmarshal raw response data
	data := &setupResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(setupAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(setupAction, *data.Msg)
	}

	return nil
}
