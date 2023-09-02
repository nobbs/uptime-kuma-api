package state_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/stretchr/testify/assert"
)

func TestState_Info(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	tests := []struct {
		name    string
		fields  fields
		want    *state.Info
		wantErr bool
	}{
		{
			"Set info",
			fields{
				s: state.NewState(),
				setup: func(s *state.State) {
					err := s.SetInfo(&state.Info{})
					assert.NoError(t, err, "Should not be error")
				},
			},
			&state.Info{},
			false,
		},
		{
			"Info is nil",
			fields{
				s: &state.State{},
			},
			nil,
			true,
		},
		{
			"state is nil",
			fields{
				s: nil,
				setup: func(s *state.State) {
					err := s.SetInfo(nil)
					assert.Error(t, err, "Should be error")
				},
			},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			got, err := tt.fields.s.Info()
			assert.Equal(t, tt.wantErr, err != nil, "Should be same error")
			assert.Equal(t, tt.want, got, "Should be same value")
		})
	}
}
