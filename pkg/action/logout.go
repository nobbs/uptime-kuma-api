package action

import (
	"fmt"

	"github.com/nobbs/uptime-kuma-api/pkg/handler"
)

const (
	logoutAction = "logout"
)

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
	c.State().SetLoggedIn(false)

	return nil
}
