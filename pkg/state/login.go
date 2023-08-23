package state

import (
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

// LoggedIn returns the login state of the client.
func (s *State) LoggedIn() (bool, error) {
	if s == nil {
		return false, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.loggedIn == nil {
		return false, ErrNotSetYet
	}

	return *s.loggedIn, nil
}

// SetLoggedIn sets the login state of the client.
func (s *State) SetLoggedIn(loggedIn bool) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.loggedIn = utils.NewBool(loggedIn)

	return nil
}
