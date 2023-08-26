package state

import (
	"github.com/go-playground/validator/v10"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

// Monitor represents a monitor object.
type Monitor struct {
	AcceptedStatuscodes      []string        `json:"accepted_statuscodes"               mapstructure:"accepted_statuscodes"     validate:"required"`
	Active                   *bool           `json:"-"                                  mapstructure:"active"`
	AuthDomain               *string         `json:"authDomain,omitempty"               mapstructure:"authDomain"`
	AuthMethod               *string         `json:"authMethod,omitempty"               mapstructure:"authMethod"`
	AuthWorkstation          *string         `json:"authWorkstation,omitempty"          mapstructure:"authWorkstation"`
	BasicAuthPass            *string         `json:"basic_auth_pass,omitempty"          mapstructure:"basic_auth_pass"`
	BasicAuthUser            *string         `json:"basic_auth_user,omitempty"          mapstructure:"basic_auth_user"`
	Body                     *string         `json:"body,omitempty"                     mapstructure:"body"`
	ChildrenIds              []int           `json:"-"                                  mapstructure:"childrenIds"`
	DatabaseConnectionString *string         `json:"databaseConnectionString,omitempty" mapstructure:"databaseConnectionString"`
	DatabaseQuery            *string         `json:"databaseQuery,omitempty"            mapstructure:"databaseQuery"`
	Description              *string         `json:"description,omitempty"              mapstructure:"description"`
	DnsLastResult            *string         `json:"dns_last_result,omitempty"          mapstructure:"dns_last_result"`
	DnsResolveServer         *string         `json:"dns_resolve_server,omitempty"       mapstructure:"dns_resolve_server"       validate:"required_if=Type dns"`
	DnsResolveType           *string         `json:"dns_resolve_type,omitempty"         mapstructure:"dns_resolve_type"`
	DockerContainer          *string         `json:"docker_container,omitempty"         mapstructure:"docker_container"         validate:"required_if=Type docker"`
	DockerHost               *string         `json:"docker_host,omitempty"              mapstructure:"docker_host"              validate:"required_if=Type docker"`
	ExpiryNotification       *bool           `json:"expiryNotification,omitempty"       mapstructure:"expiryNotification"`
	ForceInactive            *bool           `json:"-"                                  mapstructure:"forceInactive"`
	Game                     *string         `json:"game,omitempty"                     mapstructure:"game"                     validate:"required_if=Type gamedig"`
	GrpcBody                 *string         `json:"grpcBody,omitempty"                 mapstructure:"grpcBody"`
	GrpcEnableTls            *bool           `json:"grpcEnableTls,omitempty"            mapstructure:"grpcEnableTls"`
	GrpcMetadata             *string         `json:"grpcMetadata,omitempty"             mapstructure:"grpcMetadata"`
	GrpcMethod               *string         `json:"grpcMethod,omitempty"               mapstructure:"grpcMethod"               validate:"required_if=Type grpc-keyword"`
	GrpcProtobuf             *string         `json:"grpcProtobuf,omitempty"             mapstructure:"grpcProtobuf"`
	GrpcServiceName          *string         `json:"grpcServiceName,omitempty"          mapstructure:"grpcServiceName"          validate:"required_if=Type grpc-keyword"`
	GrpcUrl                  *string         `json:"grpcUrl,omitempty"                  mapstructure:"grpcUrl"                  validate:"required_if=Type grpc-keyword"`
	Headers                  []string        `json:"headers,omitempty"                  mapstructure:"headers"`
	Hostname                 *string         `json:"hostname,omitempty"                 mapstructure:"hostname"                 validate:"required_if=Type port,required_if=Type ping,required_if=Type dns,required_if=Type steam,required_if=Type gamedig,required_if=Type mqtt"`
	HttpBodyEncoding         *string         `json:"httpBodyEncoding,omitempty"         mapstructure:"httpBodyEncoding"`
	Id                       *int            `json:"-"                                  mapstructure:"id"`
	IgnoreTls                *bool           `json:"ignoreTls,omitempty"                mapstructure:"ignoreTls"`
	IncludeSensitiveData     *bool           `json:"-"                                  mapstructure:"includeSensitiveData"`
	Interval                 *int            `json:"interval"                           mapstructure:"interval"                 validate:"required,min=20"`
	InvertKeyword            *bool           `json:"invertKeyword,omitempty"            mapstructure:"invertKeyword"`
	Keyword                  *string         `json:"keyword,omitempty"                  mapstructure:"keyword"                  validate:"required_if=Type keyword,required_if=Type grpc-keyword"`
	KeywordType              *string         `json:"-"                                  mapstructure:"keywordType"`
	Maintenance              *bool           `json:"-"                                  mapstructure:"maintenance"`
	Maxredirects             *int            `json:"maxredirects,omitempty"             mapstructure:"maxredirects"             validate:"min=0,required_if=Type http,required_if=Type keyword"`
	Maxretries               *int            `json:"maxretries"                         mapstructure:"maxretries"               validate:"min=0"`
	Method                   *string         `json:"method,omitempty"                   mapstructure:"method"`
	MqttPassword             *string         `json:"mqttPassword,omitempty"             mapstructure:"mqttPassword"`
	MqttSuccessMessage       *string         `json:"mqttSuccessMessage,omitempty"       mapstructure:"mqttSuccessMessage"`
	MqttTopic                *string         `json:"mqttTopic,omitempty"                mapstructure:"mqttTopic"                validate:"required_if=Type mqtt"`
	MqttUsername             *string         `json:"mqttUsername,omitempty"             mapstructure:"mqttUsername"`
	Name                     *string         `json:"name"                               mapstructure:"name"                     validate:"required"`
	NotificationIDList       map[string]bool `json:"notificationIDList,omitempty"       mapstructure:"notificationIDList"`
	PacketSize               *int            `json:"packetSize,omitempty"               mapstructure:"packetSize"`
	Parent                   *int            `json:"parent,omitempty"                   mapstructure:"parent"`
	PathName                 *string         `json:"-"                                  mapstructure:"pathName"`
	Port                     *int            `json:"port"                               mapstructure:"port"                     validate:"omitempty,hostname_port,omitempty,required_if=Type port,required_if=Type dns,required_if=Type steam,required_if=Type gamedig,required_if=Type mqtt"`
	ProxyId                  *int            `json:"proxyId,omitempty"                  mapstructure:"proxyId"`
	PushToken                *string         `json:"pushToken,omitempty"                mapstructure:"pushToken"`
	RadiusCalledStationId    *string         `json:"radiusCalledStationId,omitempty"    mapstructure:"radiusCalledStationId"    validate:"required_if=Type radius"`
	RadiusCallingStationId   *string         `json:"radiusCallingStationId,omitempty"   mapstructure:"radiusCallingStationId"   validate:"required_if=Type radius"`
	RadiusPassword           *string         `json:"radiusPassword,omitempty"           mapstructure:"radiusPassword"           validate:"required_if=Type radius"`
	RadiusSecret             *string         `json:"radiusSecret,omitempty"             mapstructure:"radiusSecret"             validate:"required_if=Type radius"`
	RadiusUsername           *string         `json:"radiusUsername,omitempty"           mapstructure:"radiusUsername"           validate:"required_if=Type radius"`
	ResendInterval           *int            `json:"resendInterval,omitempty"           mapstructure:"resendInterval"`
	RetryInterval            *int            `json:"retryInterval"                      mapstructure:"retryInterval"            validate:"required,min=20"`
	Tags                     []Tag           `json:"-"                                  mapstructure:"tags"`
	TlsCa                    *string         `json:"tlsCa,omitempty"                    mapstructure:"tlsCa"`
	TlsCert                  *string         `json:"tlsCert,omitempty"                  mapstructure:"tlsCert"`
	TlsKey                   *string         `json:"tlsKey,omitempty"                   mapstructure:"tlsKey"`
	Type                     *string         `json:"type"                               mapstructure:"type"                     validate:"required,monitorType"`
	UpsideDown               *bool           `json:"upsideDown,omitempty"               mapstructure:"upsideDown"`
	Url                      *string         `json:"url,omitempty"                      mapstructure:"url"                      validate:"required_if=Type http,required_if=Type keyword"`
	Weight                   *int            `json:"weight,omitempty"                   mapstructure:"weight"`

	// additional fields that may be present but are not documented
	Unmapped map[string]any `json:"-" mapstructure:",remain"`
}

// init registers the monitor type validator.
func init() {
	err := utils.GetValidator().RegisterValidation("monitorType", validateMonitorType)
	if err != nil {
		panic(err)
	}
}

// validateMonitorType validates the monitor type.
func validateMonitorType(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "dns", "docker", "gamedig", "group", "grpc-keyword", "http", "json-query", "kafka-producer", "keyword", "mongodb", "mqtt", "mysql", "ping", "port", "postgres", "push", "radius", "redis", "sqlServer", "steam":
		return true
	default:
		return false
	}
}

// Validate validates the monitor.
func (m *Monitor) Validate() error {
	return utils.ValidateStruct(m)
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
