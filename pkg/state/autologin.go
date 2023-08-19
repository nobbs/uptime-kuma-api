package state

import "github.com/nobbs/uptime-kuma-api/pkg/utils"

// AutoLogin returns the auto login state of the client.
func (s *State) AutoLogin() (bool, error) {
	if s == nil {
		return false, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.autoLogin == nil {
		return false, ErrNotSetYet
	}

	return *s.autoLogin, nil
}

// SetAutoLogin sets the auto login state of the client.
func (s *State) SetAutoLogin(autoLogin bool) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.autoLogin = utils.NewBool(autoLogin)
	return nil
}
