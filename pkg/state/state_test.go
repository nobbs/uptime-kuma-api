package state_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

func TestState_MarkSeen(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	type args struct {
		event string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"State nil",
			fields{
				s: nil,
			},
			args{
				event: "test",
			},
		},
		{
			"Mark seen",
			fields{
				s: state.NewState(),
			},
			args{
				event: "test",
			},
		},
		{
			"Mark seen twice",
			fields{
				s: state.NewState(),
			},
			args{
				event: "test",
			},
		},
		{
			"Mark seen with empty event",
			fields{
				s: state.NewState(),
			},
			args{
				event: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			tt.fields.s.MarkSeen(tt.args.event)
		})
	}
}

func TestState_HasSeen(t *testing.T) {
	type fields struct {
		s     *state.State
		setup func(*state.State)
	}

	type args struct {
		event string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"State nil",
			fields{
				s: nil,
			},
			args{
				event: "test",
			},
			false,
		},
		{
			"Has seen",
			fields{
				s: state.NewState(),
				setup: func(s *state.State) {
					s.MarkSeen("test")
				},
			},
			args{
				event: "test",
			},
			true,
		},
		{
			"Has not seen",
			fields{
				s: state.NewState(),
			},
			args{
				event: "test",
			},
			false,
		},
		{
			"Has seen with empty event",
			fields{
				s: state.NewState(),
			},
			args{
				event: "",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.s)
			}

			got := tt.fields.s.HasSeen(tt.args.event)

			if got != tt.want {
				t.Errorf("State.HasSeen() = %v, want %v", got, tt.want)
			}
		})
	}
}
