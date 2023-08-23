package state

import (
	"sync"

	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

// State stores the data received from Uptime Kuma. It is used both to store the
// data received from the server either through events or actions.
type State struct {
	mu sync.RWMutex

	// Store all occurred events.
	seenEvents *SeenEvents

	// Stores the connection state of the client.
	connected *bool

	// Stores the login state of the client.
	loggedIn *bool

	// Stores the auto login state of the client.
	autoLogin *bool

	// Stores the info data.
	info *Info

	// Stores the heartbeats.
	heartbeats map[int]HeartbeatQueue

	// Stores the important heartbeats.
	importantHeartbeats map[int]HeartbeatQueue

	// Stores the tags
	tags map[int]*Tag
}

// NewState creates a new empty state instance.
func NewState() *State {
	return &State{
		seenEvents:          &SeenEvents{},
		connected:           nil,
		loggedIn:            utils.NewBool(false),
		autoLogin:           nil,
		info:                nil,
		heartbeats:          nil,
		importantHeartbeats: nil,
		tags:                nil,
	}
}

// HasSeen returns true if the given event has occurred.
func (s *State) HasSeen(event string) bool {
	if s == nil {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.seenEvents.HasSeen(event)
}

// MarkSeen marks the given event as occurred.
func (s *State) MarkSeen(event string) {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.seenEvents.MarkSeen(event)
}
