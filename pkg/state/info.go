package state

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
		return nil, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.info == nil {
		return nil, ErrNotSetYet
	}

	return s.info, nil
}

// SetInfo sets the info data received from Uptime Kuma.
func (s *State) SetInfo(info *Info) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.info = info

	return nil
}
