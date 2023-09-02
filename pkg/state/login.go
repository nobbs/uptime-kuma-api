package state

import (
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/nobbs/uptime-kuma-api/pkg/xerrors"
)

// LoggedIn returns the login state of the client.
func (s *State) LoggedIn() (bool, error) {
	if s == nil {
		return false, xerrors.ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.loggedIn == nil {
		return false, xerrors.ErrNotSetYet
	}

	return *s.loggedIn, nil
}

// SetLoggedIn sets the login state of the client.
func (s *State) SetLoggedIn(loggedIn bool) error {
	if s == nil {
		return xerrors.ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.loggedIn = utils.NewBool(loggedIn)

	return nil
}
