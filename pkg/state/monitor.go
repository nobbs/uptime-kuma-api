package state

// Monitor represents a monitor object.
type Monitor struct {
	Active        bool   `mapstructure:"active"`
	Interval      int    `mapstructure:"interval"`
	Maxretries    int    `mapstructure:"maxretries"`
	Name          string `mapstructure:"name"`
	RetryInterval int    `mapstructure:"retryInterval"`
	Type          string `mapstructure:"type"`

	AcceptedStatuscodes      []string       `mapstructure:"accepted_statuscodes"`
	AuthDomain               *string        `mapstructure:"authDomain"`
	AuthMethod               *string        `mapstructure:"authMethod"`
	AuthWorkstation          *string        `mapstructure:"authWorkstation"`
	BasicAuthPass            *string        `mapstructure:"basic_auth_pass"`
	BasicAuthUser            *string        `mapstructure:"basic_auth_user"`
	Body                     *string        `mapstructure:"body"`
	ChildrenIds              []int          `mapstructure:"childrenIds"`
	DatabaseConnectionString *string        `mapstructure:"databaseConnectionString"`
	DatabaseQuery            *string        `mapstructure:"databaseQuery"`
	Description              *string        `mapstructure:"description"`
	DnsLastResult            *string        `mapstructure:"dns_last_result"`
	DnsResolveServer         *string        `mapstructure:"dns_resolve_server"`
	DnsResolveType           *string        `mapstructure:"dns_resolve_type"`
	DockerContainer          *string        `mapstructure:"docker_container"`
	DockerHost               *string        `mapstructure:"docker_host"`
	ExpiryNotification       *bool          `mapstructure:"expiryNotification"`
	ForceInactive            *bool          `mapstructure:"forceInactive"`
	Game                     *string        `mapstructure:"game"`
	GrpcBody                 *string        `mapstructure:"grpcBody"`
	GrpcEnableTls            *bool          `mapstructure:"grpcEnableTls"`
	GrpcMetadata             *string        `mapstructure:"grpcMetadata"`
	GrpcMethod               *string        `mapstructure:"grpcMethod"`
	GrpcProtobuf             *string        `mapstructure:"grpcProtobuf"`
	GrpcServiceName          *string        `mapstructure:"grpcServiceName"`
	GrpcUrl                  *string        `mapstructure:"grpcUrl"`
	Headers                  []string       `mapstructure:"headers"`
	Hostname                 *string        `mapstructure:"hostname"`
	HttpBodyEncoding         *string        `mapstructure:"httpBodyEncoding"`
	Id                       *int           `mapstructure:"id"`
	IgnoreTls                *bool          `mapstructure:"ignoreTls"`
	IncludeSensitiveData     *bool          `mapstructure:"includeSensitiveData"`
	InvertKeyword            *bool          `mapstructure:"invertKeyword"`
	Keyword                  *string        `mapstructure:"keyword"`
	KeywordType              *string        `mapstructure:"keywordType"` // TODO?
	Maintenance              *bool          `mapstructure:"maintenance"`
	Maxredirects             *int           `mapstructure:"maxredirects"`
	Method                   *string        `mapstructure:"method"`
	MqttPassword             *string        `mapstructure:"mqttPassword"`
	MqttSuccessMessage       *string        `mapstructure:"mqttSuccessMessage"`
	MqttTopic                *string        `mapstructure:"mqttTopic"`
	MqttUsername             *string        `mapstructure:"mqttUsername"`
	NotificationIDList       map[int]string `mapstructure:"notificationIDList"`
	PacketSize               *int           `mapstructure:"packetSize"`
	Parent                   *int           `mapstructure:"parent"`
	PathName                 *string        `mapstructure:"pathName"`
	Port                     *int           `mapstructure:"port"`
	ProxyId                  *int           `mapstructure:"proxyId"`
	PushToken                *string        `mapstructure:"pushToken"`
	RadiusCalledStationId    *string        `mapstructure:"radiusCalledStationId"`
	RadiusCallingStationId   *string        `mapstructure:"radiusCallingStationId"`
	RadiusPassword           *string        `mapstructure:"radiusPassword"`
	RadiusSecret             *string        `mapstructure:"radiusSecret"`
	RadiusUsername           *string        `mapstructure:"radiusUsername"`
	ResendInterval           *int           `mapstructure:"resendInterval"`
	Tags                     []string       `mapstructure:"tags"`
	TlsCa                    *string        `mapstructure:"tlsCa"`
	TlsCert                  *string        `mapstructure:"tlsCert"`
	TlsKey                   *string        `mapstructure:"tlsKey"`
	UpsideDown               *bool          `mapstructure:"upsideDown"`
	Url                      *string        `mapstructure:"url"`
	Weight                   *int           `mapstructure:"weight"`

	// additional fields that may be present but are not documented
	Unmapped map[string]any `mapstructure:",remain"`
}

// Monitors returns all monitors received from Uptime Kuma.
func (s *State) Monitors() (map[int]*Monitor, error) {
	if s == nil {
		return nil, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.monitors == nil {
		return nil, ErrNotSetYet
	}

	return s.monitors, nil
}

// Monitor returns the monitor with the given id.
func (s *State) Monitor(monitorId int) (*Monitor, error) {
	if s == nil {
		return nil, ErrStateNil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.monitors == nil {
		return nil, ErrNotSetYet
	}

	monitor, ok := s.monitors[monitorId]
	if !ok {
		return nil, NewErrNotFound("monitor", monitorId)
	}

	return monitor, nil
}

// SetMonitors sets the monitors received from Uptime Kuma.
func (s *State) SetMonitors(monitors map[int]*Monitor) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.monitors = monitors

	return nil
}

// SetMonitor sets the monitor with the given id.
func (s *State) SetMonitor(id int, monitor *Monitor) error {
	if s == nil {
		return ErrStateNil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.monitors == nil {
		s.monitors = make(map[int]*Monitor)
	}

	s.monitors[id] = monitor

	return nil
}
