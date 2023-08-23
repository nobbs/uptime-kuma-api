package action

import (
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
)

const (
	getSettingsAction = "getSettings"
	setSettingsAction = "setSettings"
)

// Settings represents the settings of the Uptime Kuma instance.
type Settings struct {
	CheckUpdate         *bool   `mapstructure:"checkUpdate" json:"checkUpdate,omitempty"`
	CheckBeta           *bool   `mapstructure:"checkBeta" json:"checkBeta,omitempty"`
	KeepDataPeriodDays  *int    `mapstructure:"keepDataPeriodDays" json:"keepDataPeriodDays,omitempty"`
	ServerTimezone      *string `mapstructure:"serverTimezone" json:"serverTimezone,omitempty"`
	EntryPage           *string `mapstructure:"entryPage" json:"entryPage,omitempty"`
	SearchEngineIndex   *bool   `mapstructure:"searchEngineIndex" json:"searchEngineIndex,omitempty"`
	PrimaryBaseURL      *string `mapstructure:"primaryBaseURL" json:"primaryBaseURL,omitempty"`
	SteamAPIKey         *string `mapstructure:"steamAPIKey" json:"steamAPIKey,omitempty"`
	DnsCache            *bool   `mapstructure:"dnsCache" json:"dnsCache,omitempty"`
	TlsExpiryNotifyDays []int   `mapstructure:"tlsExpiryNotifyDays" json:"tlsExpiryNotifyDays,omitempty"`
	DisableAuth         *bool   `mapstructure:"disableAuth" json:"disableAuth,omitempty"`
	TrustProxy          *bool   `mapstructure:"trustProxy" json:"trustProxy,omitempty"`

	Unmapped map[string]any `mapstructure:",remain"`
}

type getSettingsResponse struct {
	Ok   bool      `mapstructure:"ok"`
	Msg  *string   `mapstructure:"msg"`
	Data *Settings `mapstructure:"data"`
}

type setSettingsResponse struct {
	Ok  bool    `mapstructure:"ok"`
	Msg *string `mapstructure:"msg"`
}

// GetSettings returns the settings of the Uptime Kuma instance.
func GetSettings(c StatefulEmiter) (*Settings, error) {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return nil, NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	reseponse, err := c.Emit(getSettingsAction, defaultEmitTimeout)
	if err != nil {
		return nil, NewErrActionFailed(getSettingsAction, err.Error())
	}

	// unmarshal raw response data
	data := &getSettingsResponse{}
	if err := decode(reseponse, data); err != nil {
		return nil, NewErrActionFailed(getSettingsAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return nil, NewErrActionFailed(getSettingsAction, *data.Msg)
	}

	return data.Data, nil
}

// SetSettings sets the settings of the Uptime Kuma instance.
func SetSettings(c StatefulEmiter, settings *Settings, password string) error {
	// ensure client is connected
	if err := c.Await(handler.ConnectEvent, defaultAwaitTimeout); err != nil {
		return NewErrAwaitFailed(handler.ConnectEvent, err)
	}

	// call action
	reseponse, err := c.Emit(setSettingsAction, defaultEmitTimeout, settings, password)
	if err != nil {
		return NewErrActionFailed(setSettingsAction, err.Error())
	}

	// unmarshal raw response data
	data := &setSettingsResponse{}
	if err := decode(reseponse, data); err != nil {
		return NewErrActionFailed(setSettingsAction, err.Error())
	}

	// check if action was successful
	if !data.Ok {
		return NewErrActionFailed(setSettingsAction, *data.Msg)
	}

	return nil
}
