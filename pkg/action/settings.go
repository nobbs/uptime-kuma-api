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
	CheckUpdate         *bool   `json:"checkUpdate,omitempty"         mapstructure:"checkUpdate"`
	CheckBeta           *bool   `json:"checkBeta,omitempty"           mapstructure:"checkBeta"`
	KeepDataPeriodDays  *int    `json:"keepDataPeriodDays,omitempty"  mapstructure:"keepDataPeriodDays"`
	ServerTimezone      *string `json:"serverTimezone,omitempty"      mapstructure:"serverTimezone"`
	EntryPage           *string `json:"entryPage,omitempty"           mapstructure:"entryPage"`
	SearchEngineIndex   *bool   `json:"searchEngineIndex,omitempty"   mapstructure:"searchEngineIndex"`
	PrimaryBaseURL      *string `json:"primaryBaseURL,omitempty"      mapstructure:"primaryBaseURL"`
	SteamAPIKey         *string `json:"steamAPIKey,omitempty"         mapstructure:"steamAPIKey"`
	DnsCache            *bool   `json:"dnsCache,omitempty"            mapstructure:"dnsCache"`
	TlsExpiryNotifyDays []int   `json:"tlsExpiryNotifyDays,omitempty" mapstructure:"tlsExpiryNotifyDays"`
	DisableAuth         *bool   `json:"disableAuth,omitempty"         mapstructure:"disableAuth"`
	TrustProxy          *bool   `json:"trustProxy,omitempty"          mapstructure:"trustProxy"`

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
