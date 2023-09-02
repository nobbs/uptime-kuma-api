package action

import (
	"fmt"

	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	. "github.com/nobbs/uptime-kuma-api/pkg/xerrors"
)

const (
	loginAction          = "login"
	loginByTokenAction   = "loginByToken"
	logoutAction         = "logout"
	changePasswordAction = "changePassword"
	prepare2faAction     = "prepare2FA"
	save2faAction        = "save2FA"
	disable2faAction     = "disable2FA"
	verifyTokenAction    = "verifyToken"
	twoFAStatusAction    = "twoFAStatus"
	needSetupAction      = "needSetup"
	setupAction          = "setup"
)

// loginRequest is the request payload for the login action.
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}

// loginResponse is the response payload for the login action.
type loginResponse struct {
	Ok            bool    `mapstructure:"ok"`
	Msg           *string `mapstructure:"msg"`
	Token         *string `mapstructure:"token"`
	TokenRequired *bool   `mapstructure:"tokenRequired"`
}

// changePasswordRequest is the request payload for the change password action.
type changePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// changePasswordResponse is the response payload for the change password action.
type changePasswordResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

// prepare2faResponse is the response payload for the prepare 2fa action.
type prepare2faResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
	Uri *string `mapstructure:"uri"`
}

// save2faResponse is the response payload for the save 2fa action.
type save2faResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

// disable2faResponse is the response payload for the disable 2fa action.
type disable2faResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

// verifyTokenResponse is the response payload for the verify token action.
type verifyTokenResponse struct {
	Ok    bool    `mapstructure:"ok"`
	Msg   *string `mapstructure:"msg"`
	Valid *bool   `mapstructure:"valid"`
}

// twoFAStatusResponse is the response payload for the 2fa status action.
type twoFAStatusResponse struct {
	Ok     bool    `mapstructure:"ok"`
	Msg    *string `mapstructure:"msg"`
	Status *bool   `mapstructure:"status"`
}

// setupResponse is the response payload for the setup action.
type setupResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

// Login sends a login request to the server with the provided username, password, and token. If the
// username and password are empty, it checks if auto login is enabled and returns an empty
// response. Returns a token and nil error if login is successful, otherwise returns an error.
func Login(c StatefulEmiter, username, password, token string) (string, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return "", NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// if username and password are empty, check if auto login is enabled
	if username == "" && password == "" {
		if err := c.Await(handler.AutoLoginEvent, defaultAwaitTimeout); err != nil {
			return "", NewErrAwaitFailed(handler.AutoLoginEvent, err)
		}

		// auto login is enabled, return empty response
		return "", nil
	}

	// otherwise, send login request
	request := &loginRequest{
		Username: username,
		Password: password,
		Token:    token,
	}

	response, err := c.Emit(loginAction, defaultEmitTimeout, request)
	if err != nil {
		return "", NewErrActionFailed(loginAction, err.Error())
	}

	// unmarshal data
	data := &loginResponse{}
	if err := decode(response, data); err != nil {
		return "", NewErrActionFailed(loginAction, err.Error())
	}

	// check if token is required
	if data.TokenRequired != nil && *data.TokenRequired {
		return "", Err2faTokenRequired
	}

	// check if login was successful
	if !data.Ok {
		return "", NewErrLoginFailed(*data.Msg)
	}

	// set logged in to true
	if err := c.State().SetLoggedIn(true); err != nil {
		return "", err
	}

	return *data.Token, nil
}

// LoginByToken logs in the client using a token. It waits for the client to connect, sends a login
// by token request, decodes the response into a struct, checks if the login was successful, and
// sets the client's logged in status to true if successful. Returns an error if any of the steps
// fail.
func LoginByToken(c StatefulEmiter, token string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// send login by token request
	response, err := c.Emit(loginByTokenAction, defaultEmitTimeout, token)
	if err != nil {
		return NewErrActionFailed(loginByTokenAction, err.Error())
	}

	// decode response into struct
	data := &loginResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(loginByTokenAction, err.Error())
	}

	// check if login was successful
	if !data.Ok {
		return NewErrLoginFailed(*data.Msg)
	}

	// set logged in to true
	if err := c.State().SetLoggedIn(true); err != nil {
		return err
	}

	return nil
}

// Log out the current user.
func Logout(c StatefulEmiter) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// send logout request to server, ignore response
	if _, err := c.Emit(logoutAction, defaultEmitTimeout); err != nil {
		return fmt.Errorf("%s: %w", logoutAction, err)
	}

	// set logged in to false
	if err := c.State().SetLoggedIn(false); err != nil {
		return fmt.Errorf("set logged in: %w", err)
	}

	return nil
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

// Prepare2FA prepares the client for 2fa by sending a request to the server and returning the
// Uri required to generate TOTP codes. Returns an error if the request fails.
func Prepare2FA(c StatefulEmiter, currentPassword string) (string, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return "", NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(prepare2faAction, defaultEmitTimeout, currentPassword)
	if err != nil {
		return "", NewErrActionFailed(prepare2faAction, err.Error())
	}

	// unmarshal raw response data
	data := &prepare2faResponse{}
	if err := decode(response, data); err != nil {
		return "", NewErrActionFailed(prepare2faAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return "", NewErrActionFailed(prepare2faAction, *data.Msg)
	}

	return *data.Uri, nil
}

// Save2FA saves the 2fa Uri to the server thus enabling 2fa for the client. Returns an error if the
// request fails.
func Save2FA(c StatefulEmiter, currentPassword string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(save2faAction, defaultEmitTimeout, currentPassword)
	if err != nil {
		return NewErrActionFailed(save2faAction, err.Error())
	}

	// unmarshal raw response data
	data := &save2faResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(save2faAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(save2faAction, *data.Msg)
	}

	return nil
}

// Disable2FA disables 2fa for the client. Returns an error if the request fails.
func Disable2FA(c StatefulEmiter, currentPassword string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(disable2faAction, defaultEmitTimeout, currentPassword)
	if err != nil {
		return NewErrActionFailed(disable2faAction, err.Error())
	}

	// unmarshal raw response data
	data := &disable2faResponse{}
	if err := decode(response, data); err != nil {
		return NewErrActionFailed(disable2faAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(disable2faAction, *data.Msg)
	}

	return nil
}

// VerifyToken verifies the provided 2fa token. Returns an error if the request fails.
func VerifyToken(c StatefulEmiter, currentPassword, token string) (bool, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return false, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(verifyTokenAction, defaultEmitTimeout, token, currentPassword)
	if err != nil {
		return false, NewErrActionFailed(verifyTokenAction, err.Error())
	}

	// unmarshal raw response data
	data := &verifyTokenResponse{}
	if err := decode(response, data); err != nil {
		return false, NewErrActionFailed(verifyTokenAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return false, NewErrActionFailed(verifyTokenAction, *data.Msg)
	}

	return *data.Valid, nil
}

// TwoFAStatus returns the current 2fa status. Returns an error if the request fails.
func TwoFAStatus(c StatefulEmiter) (bool, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return false, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	response, err := c.Emit(twoFAStatusAction, defaultEmitTimeout)
	if err != nil {
		return false, NewErrActionFailed(twoFAStatusAction, err.Error())
	}

	// unmarshal raw response data
	data := &twoFAStatusResponse{}
	if err := decode(response, data); err != nil {
		return false, NewErrActionFailed(twoFAStatusAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return false, NewErrActionFailed(twoFAStatusAction, *data.Msg)
	}

	return *data.Status, nil
}

// NeedSetup returns true if the server needs to be setup, otherwise returns false. Returns an
// error if the request fails.
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

// Setup sets up the server with the provided username and password. Returns an error if the request
// fails.
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
