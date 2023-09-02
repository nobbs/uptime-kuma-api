package state

import (
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/nobbs/uptime-kuma-api/pkg/xerrors"
)

// AutoLogin returns the auto login state of the client.
func (s *State) AutoLogin() (bool, error) {
	if s == nil {
		return false, xerrors.ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.autoLogin == nil {
		return false, xerrors.ErrNotSetYet
	}

	return *s.autoLogin, nil
}

// SetAutoLogin sets the auto login state of the client.
func (s *State) SetAutoLogin(autoLogin bool) error {
	if s == nil {
		return xerrors.ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.autoLogin = utils.NewBool(autoLogin)

	return nil
}
