package state

// Monitor represents a monitor object.
type Monitor struct {
	AcceptedStatuscodes      []string       `mapstructure:"accepted_statuscodes" json:"accepted_statuscodes"`
	Active                   bool           `mapstructure:"active" json:"-"`
	AuthDomain               *string        `mapstructure:"authDomain" json:"authDomain"`
	AuthMethod               *string        `mapstructure:"authMethod" json:"authMethod"`
	AuthWorkstation          *string        `mapstructure:"authWorkstation" json:"authWorkstation"`
	BasicAuthPass            *string        `mapstructure:"basic_auth_pass" json:"basic_auth_pass"`
	BasicAuthUser            *string        `mapstructure:"basic_auth_user" json:"basic_auth_user"`
	Body                     *string        `mapstructure:"body" json:"body"`
	ChildrenIds              []int          `mapstructure:"childrenIds" json:"-"`
	DatabaseConnectionString *string        `mapstructure:"databaseConnectionString" json:"databaseConnectionString"`
	DatabaseQuery            *string        `mapstructure:"databaseQuery" json:"databaseQuery"`
	Description              *string        `mapstructure:"description" json:"description"`
	DnsLastResult            *string        `mapstructure:"dns_last_result" json:"dns_last_result"`
	DnsResolveServer         *string        `mapstructure:"dns_resolve_server" json:"dns_resolve_server"`
	DnsResolveType           *string        `mapstructure:"dns_resolve_type" json:"dns_resolve_type"`
	DockerContainer          *string        `mapstructure:"docker_container" json:"docker_container"`
	DockerHost               *string        `mapstructure:"docker_host" json:"docker_host"`
	ExpiryNotification       *bool          `mapstructure:"expiryNotification" json:"expiryNotification"`
	ForceInactive            *bool          `mapstructure:"forceInactive" json:"-"`
	Game                     *string        `mapstructure:"game" json:"game"`
	GrpcBody                 *string        `mapstructure:"grpcBody" json:"grpcBody"`
	GrpcEnableTls            bool           `mapstructure:"grpcEnableTls" json:"grpcEnableTls"`
	GrpcMetadata             *string        `mapstructure:"grpcMetadata" json:"grpcMetadata"`
	GrpcMethod               *string        `mapstructure:"grpcMethod" json:"grpcMethod"`
	GrpcProtobuf             *string        `mapstructure:"grpcProtobuf" json:"grpcProtobuf"`
	GrpcServiceName          *string        `mapstructure:"grpcServiceName" json:"grpcServiceName"`
	GrpcUrl                  *string        `mapstructure:"grpcUrl" json:"grpcUrl"`
	Headers                  []string       `mapstructure:"headers" json:"headers"`
	Hostname                 *string        `mapstructure:"hostname" json:"hostname"`
	HttpBodyEncoding         *string        `mapstructure:"httpBodyEncoding" json:"httpBodyEncoding"`
	Id                       int            `mapstructure:"id" json:"-"`
	IgnoreTls                bool           `mapstructure:"ignoreTls" json:"ignoreTls"`
	IncludeSensitiveData     *bool          `mapstructure:"includeSensitiveData" json:"-"`
	Interval                 int            `mapstructure:"interval" json:"interval"`
	InvertKeyword            bool           `mapstructure:"invertKeyword" json:"invertKeyword"`
	Keyword                  *string        `mapstructure:"keyword" json:"keyword"`
	KeywordType              *string        `mapstructure:"keywordType" json:"-"`
	Maintenance              *bool          `mapstructure:"maintenance" json:"-"`
	Maxredirects             int            `mapstructure:"maxredirects" json:"maxredirects"`
	Maxretries               int            `mapstructure:"maxretries" json:"maxretries"`
	Method                   *string        `mapstructure:"method" json:"method"`
	MqttPassword             *string        `mapstructure:"mqttPassword" json:"mqttPassword"`
	MqttSuccessMessage       *string        `mapstructure:"mqttSuccessMessage" json:"mqttSuccessMessage"`
	MqttTopic                *string        `mapstructure:"mqttTopic" json:"mqttTopic"`
	MqttUsername             *string        `mapstructure:"mqttUsername" json:"mqttUsername"`
	Name                     string         `mapstructure:"name" json:"name"`
	NotificationIDList       map[int]string `mapstructure:"notificationIDList" json:"notificationIDList"`
	PacketSize               int            `mapstructure:"packetSize" json:"packetSize"`
	Parent                   *int           `mapstructure:"parent" json:"parent"`
	PathName                 *string        `mapstructure:"pathName" json:"-"`
	Port                     *int           `mapstructure:"port" json:"port"`
	ProxyId                  *int           `mapstructure:"proxyId" json:"proxyId"`
	PushToken                *string        `mapstructure:"pushToken" json:"pushToken"`
	RadiusCalledStationId    *string        `mapstructure:"radiusCalledStationId" json:"radiusCalledStationId"`
	RadiusCallingStationId   *string        `mapstructure:"radiusCallingStationId" json:"radiusCallingStationId"`
	RadiusPassword           *string        `mapstructure:"radiusPassword" json:"radiusPassword"`
	RadiusSecret             *string        `mapstructure:"radiusSecret" json:"radiusSecret"`
	RadiusUsername           *string        `mapstructure:"radiusUsername" json:"radiusUsername"`
	ResendInterval           int            `mapstructure:"resendInterval" json:"resendInterval"`
	RetryInterval            int            `mapstructure:"retryInterval" json:"retryInterval"`
	Tags                     []string       `mapstructure:"tags" json:"-"`
	TlsCa                    *string        `mapstructure:"tlsCa" json:"tlsCa"`
	TlsCert                  *string        `mapstructure:"tlsCert" json:"tlsCert"`
	TlsKey                   *string        `mapstructure:"tlsKey" json:"tlsKey"`
	Type                     string         `mapstructure:"type" json:"type"`
	UpsideDown               bool           `mapstructure:"upsideDown" json:"upsideDown"`
	Url                      *string        `mapstructure:"url" json:"url"`
	Weight                   *int           `mapstructure:"weight" json:"weight"`

	// additional fields that may be present but are not documented
	Unmapped map[string]any `mapstructure:",remain" json:"-"`
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
