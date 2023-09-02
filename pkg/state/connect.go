package state

import (
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/nobbs/uptime-kuma-api/pkg/xerrors"
)

// Connected returns the connection state of the client.
func (s *State) Connected() (bool, error) {
	if s == nil {
		return false, xerrors.ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.connected == nil {
		return false, xerrors.ErrNotSetYet
	}

	return *s.connected, nil
}

// SetConnected sets the connection state of the client.
func (s *State) SetConnected(connected bool) error {
	if s == nil {
		return xerrors.ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.connected = utils.NewBool(connected)

	return nil
}
