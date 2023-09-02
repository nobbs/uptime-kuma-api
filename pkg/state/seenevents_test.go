package state_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

func TestSeenEvents_MarkSeen(t *testing.T) {
	type fields struct {
		oe    state.SeenEvents
		setup func(state.SeenEvents)
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
			"SeenEvents nil",
			fields{
				oe: nil,
			},
			args{
				event: "test",
			},
		},
		{
			"Mark seen",
			fields{
				oe: state.SeenEvents{},
			},
			args{
				event: "test",
			},
		},
		{
			"Mark seen twice",
			fields{
				oe: state.SeenEvents{},
			},
			args{
				event: "test",
			},
		},
		{
			"Mark seen with empty event",
			fields{
				oe: state.SeenEvents{},
			},
			args{
				event: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.setup != nil {
				tt.fields.setup(tt.fields.oe)
			}
			tt.fields.oe.MarkSeen(tt.args.event)
		})
	}
}

func TestSeenEvents_HasSeen(t *testing.T) {
	type fields struct {
		oe    state.SeenEvents
		setup func(state.SeenEvents)
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
			"SeenEvents nil",
			fields{
				oe: nil,
			},
			args{
				event: "test",
			},
			false,
		},
		{
			"Has seen",
			fields{
				oe: state.SeenEvents{
					"test": struct{}{},
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
				oe: state.SeenEvents{},
			},
			args{
				event: "test",
			},
			false,
		},
		{
			"Has seen with empty event",
			fields{
				oe: state.SeenEvents{},
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
				tt.fields.setup(tt.fields.oe)
			}
			if got := tt.fields.oe.HasSeen(tt.args.event); got != tt.want {
				t.Errorf("SeenEvents.HasSeen() = %v, want %v", got, tt.want)
			}
		})
	}
}
