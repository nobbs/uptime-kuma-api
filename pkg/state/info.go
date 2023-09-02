package state

import "github.com/nobbs/uptime-kuma-api/pkg/xerrors"

// Info stores the information data received from Uptime Kuma.
type Info struct {
	LatestVersion        *string `mapstructure:"latestVersion"`
	PrimaryBaseURL       *string `mapstructure:"primaryBaseURL"`
	ServerTimezone       *string `mapstructure:"serverTimezone"`
	ServerTimezoneOffset *string `mapstructure:"serverTimezoneOffset"`
	Version              *string `mapstructure:"version"`
}

// Info return the info data received from Uptime Kuma.
func (s *State) Info() (*Info, error) {
	if s == nil {
		return nil, xerrors.ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.info == nil {
		return nil, xerrors.ErrNotSetYet
	}

	return s.info, nil
}

// SetInfo sets the info data received from Uptime Kuma.
func (s *State) SetInfo(info *Info) error {
	if s == nil {
		return xerrors.ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.info = info

	return nil
}
