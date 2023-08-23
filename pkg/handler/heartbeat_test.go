package handler_test

import (
	"testing"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/mocks"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/nobbs/uptime-kuma-api/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHeartbeat_Event(t *testing.T) {
	c := handler.NewHeartbeat(nil)

	assert.Equal(t, handler.HeartbeatEvent, c.Event())
}

func TestHeartbeatList_Event(t *testing.T) {
	c := handler.NewHeartbeatList(nil)

	assert.Equal(t, handler.HeartbeatListEvent, c.Event())
}

func TestImportantHeartbeatList_Event(t *testing.T) {
	c := handler.NewImportantHeartbeatList(nil)

	assert.Equal(t, handler.ImportantHeartbeatListEvent, c.Event())
}

func TestHeartbeat_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewHeartbeat(nil)

	r.EXPECT().On(handler.HeartbeatEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestHeartbeatList_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewHeartbeatList(nil)

	r.EXPECT().On(handler.HeartbeatListEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestHeartbeatImportantList_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewImportantHeartbeatList(nil)

	r.EXPECT().On(handler.ImportantHeartbeatListEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestHeartbeatList_Occurred(t *testing.T) {
	s := mocks.NewHeartbeatListState(t)
	c := handler.NewHeartbeatList(s)

	s.EXPECT().Heartbeats(0).Return(nil, state.ErrNotSetYet).Once()
	s.EXPECT().Heartbeats(0).Return([]state.Heartbeat{}, nil).Once()

	assert.False(t, c.Occurred())
	assert.True(t, c.Occurred())
}

func TestHeartbeatList_Callback(t *testing.T) {
	type fields struct {
		state *mocks.HeartbeatListState
	}

	type args struct {
		ch        *shadiaosocketio.Channel
		id        any
		result    []any
		overwrite any
	}

	tests := []struct {
		name   string
		fields *fields
		args   *args
		want   *string

		on     func(*fields)
		assert func(*testing.T, *fields)
	}{
		{
			name: "ok",
			fields: &fields{
				state: mocks.NewHeartbeatListState(t),
			},
			args: &args{
				ch: &shadiaosocketio.Channel{},
				id: 0,
				result: []any{
					map[string]any{
						"down_count": 0,
						"duration":   24,
						"id":         0,
						"important":  false,
						"monitorid":  0,
						"msg":        "none",
						"ping":       111,
						"status":     1,
						"time":       "2021-01-01T00:00:00Z",
					},
				},
				overwrite: false,
			},
			want: nil,
			on: func(f *fields) {
				f.state.EXPECT().SetHeartbeats(0, []state.Heartbeat{
					{
						DownCount: 0,
						Duration:  24,
						Id:        0,
						Important: false,
						MonitorId: 0,
						Msg:       "none",
						Ping:      111,
						Status:    true,
						Time:      "2021-01-01T00:00:00Z",
					},
				}, false).Return(nil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
		{
			name: "decode failed",
			fields: &fields{
				state: mocks.NewHeartbeatListState(t),
			},
			args: &args{
				ch:        &shadiaosocketio.Channel{},
				id:        0,
				result:    []any{},
				overwrite: [1]int{1},
			},
			want: utils.NewString("decode failed"),
			on:   func(f *fields) {},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertNotCalled(t, "SetHeartbeats", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			name: "data decode failed",
			fields: &fields{
				state: mocks.NewHeartbeatListState(t),
			},
			args: &args{
				ch: &shadiaosocketio.Channel{},
				id: 0,
				result: []any{
					"invalid result",
				},
				overwrite: false,
			},
			want: utils.NewString("decode failed"),
			on:   func(f *fields) {},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertNotCalled(t, "SetHeartbeats", mock.Anything, mock.Anything, mock.Anything)
			},
		},
		{
			name: "set heartbeat failed",
			fields: &fields{
				state: mocks.NewHeartbeatListState(t),
			},
			args: &args{
				ch:        &shadiaosocketio.Channel{},
				id:        0,
				result:    []any{},
				overwrite: false,
			},
			want: utils.NewString(state.ErrStateNil.Error()),
			on: func(f *fields) {
				f.state.EXPECT().SetHeartbeats(0, []state.Heartbeat{}, false).Return(state.ErrStateNil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup mocks
			c := handler.NewHeartbeatList(tt.fields.state)

			if tt.on != nil {
				tt.on(tt.fields)
			}

			// run function
			got := c.Callback(tt.args.ch, tt.args.id, tt.args.result, tt.args.overwrite)

			// assert results
			if tt.want != nil && assert.NotNil(t, got) {
				assert.ErrorContains(t, got, *tt.want)
			}

			if tt.assert != nil {
				tt.assert(t, tt.fields)
			}
		})
	}
}
