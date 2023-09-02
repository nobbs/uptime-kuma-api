package state_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/stretchr/testify/assert"
)

func TestState_Connected(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			"connected is true",
			fields{
				s: state.NewState(),
				setup: func(s *state.State) {
					err := s.SetConnected(true)
					assert.NoError(t, err, "Should not be error")
				},
			},
			true,
			false,
		},
		{
			"connected is false",
			fields{
				s: state.NewState(),
				setup: func(s *state.State) {
					err := s.SetConnected(false)
					assert.NoError(t, err, "Should not be error")
				},
			},
			false,
			false,
		},
		{
			"connected is nil",
			fields{
				s: state.NewState(),
			},
			false,
			true,
		},
		{
			"state is nil",
			fields{
				s: nil,
				setup: func(s *state.State) {
					err := s.SetConnected(true)
					assert.Error(t, err, "Should be error")
				},
			},
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			got, err := tt.fields.s.Connected()
			assert.Equal(t, tt.wantErr, err != nil, "Should be same error")
			assert.Equal(t, tt.want, got, "Should be same value")
		})
	}
}
