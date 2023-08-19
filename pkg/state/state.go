package state

import (
	"sync"

	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

// State stores the data received from Uptime Kuma. It is used both to store the
// data received from the server either through events or actions.
type State struct {
	mu sync.RWMutex

	// Stores the connection state of the client.
	connected *bool

	// Stores the login state of the client.
	loggedIn *bool

	// Stores the auto login state of the client.
	autoLogin *bool

	// Stores the info handler payload.
	info *Info
}

// NewState creates a new empty state instance.
func NewState() *State {
	return &State{
		connected: nil,
		loggedIn:  utils.NewBool(false),
		autoLogin: nil,
		info:      nil,
	}
}
