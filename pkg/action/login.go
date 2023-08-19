package action

import (
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
)

const (
	loginAction        = "login"
	loginByTokenAction = "loginByToken"
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
	err = decode(response, data)
	if err != nil {
		return "", NewErrActionFailed(loginAction, err.Error())
	}

	// check if token is required
	if data.TokenRequired != nil && *data.TokenRequired {
		return "", Err2faTokenRequired
	}

	// check if login was successful
	if !data.Ok {
		return "", NewErrActionFailed(loginAction, *data.Msg)
	}

	// set logged in to true
	c.State().SetLoggedIn(true)

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
	err = decode(response, data)
	if err != nil {
		return NewErrActionFailed(loginByTokenAction, err.Error())
	}

	// check if login was successful
	if !data.Ok {
		return NewErrLoginFailed(*data.Msg)
	}

	// set logged in to true
	c.State().SetLoggedIn(true)

	return nil
}
