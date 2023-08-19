package action

import (
	"encoding/json"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	// defaultEmitTimeout is the default timeout for emited events.
	defaultEmitTimeout = time.Duration(5) * time.Second

	// defaultAwaitTimeout is the default timeout for awaited events.
	defaultAwaitTimeout = time.Duration(5) * time.Second
)

// StatefulEmiter is the interface that provides the basic client methods required by the actions.
type StatefulEmiter interface {
	// Emit sends an event with the given data to the server and waits for an acknowledgement.
	Emit(string, time.Duration, ...any) (any, error)

	// Await waits for the first event with the given name to be received. If the timeout is reached
	// before the event is received, an error is returned.
	Await(string, time.Duration) error

	// State returns the current state of the client.
	State() *state.State
}

// decode decodes the raw action response into the given result struct.
func decode(raw any, result any) error {
	// assert that result is a slice of interfaces, that we only have one element
	// and that the first element is a byte slice
	r, ok := raw.([]any)
	if !ok {
		return ErrInvalidResponse
	} else if len(r) != 1 {
		return ErrInvalidResponse
	}

	rawBytes, ok := r[0].([]byte)
	if !ok {
		return ErrInvalidResponse
	}

	// unmarshal the response into a generic interface so we can properly decode it
	// into the response struct
	var rawMap any
	err := json.Unmarshal(rawBytes, &rawMap)
	if err != nil {
		return err
	}

	// unmarshal raw map into the result struct
	err = mapstructure.WeakDecode(rawMap, result)
	if err != nil {
		return err
	}

	return nil
}
