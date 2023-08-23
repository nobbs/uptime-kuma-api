package action

import (
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
)

const (
	changePasswordAction = "changePassword"
)

type changePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type changePasswordResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

func ChangePassword(c StatefulEmiter, currentPassword, newPassword string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	request := &changePasswordRequest{
		CurrentPassword: currentPassword,
		NewPassword:     newPassword,
	}

	// call action
	reseponse, err := c.Emit(changePasswordAction, defaultEmitTimeout, request)
	if err != nil {
		return NewErrActionFailed(changePasswordAction, err.Error())
	}

	// unmarshal raw response data
	data := &changePasswordResponse{}
	if err := decode(reseponse, data); err != nil {
		return NewErrActionFailed(changePasswordAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(changePasswordAction, *data.Msg)
	}

	return nil
}
