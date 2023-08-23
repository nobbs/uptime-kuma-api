package handler_test

import (
	"testing"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/gorilla/websocket"
	"github.com/nobbs/uptime-kuma-api/mocks"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDisconnect_Event(t *testing.T) {
	c := handler.NewDisconnect(nil)

	assert.Equal(t, handler.DisconnectEvent, c.Event())
}

func TestDisconnect_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewDisconnect(nil)

	r.EXPECT().On(handler.DisconnectEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestDisconnect_Occurred(t *testing.T) {
	s := mocks.NewDisconnectState(t)
	c := handler.NewDisconnect(s)

	assert.False(t, c.Occurred())
}

func TestDisconnect_Callback(t *testing.T) {
	type fields struct {
		state *mocks.DisconnectState
	}

	type args struct {
		ch     *shadiaosocketio.Channel
		reason websocket.CloseError
	}

	tests := []struct {
		name   string
		fields *fields
		args   *args
		want   error

		on     func(*fields)
		assert func(*testing.T, *fields)
	}{
		{
			name: "ok",
			fields: &fields{
				state: mocks.NewDisconnectState(t),
			},
			args: &args{
				ch:     new(shadiaosocketio.Channel),
				reason: websocket.CloseError{},
			},
			want: nil,
			on: func(f *fields) {
				f.state.EXPECT().SetConnected(false).Return(nil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
		{
			name: "nil channel",
			fields: &fields{
				state: mocks.NewDisconnectState(t),
			},
			args: &args{
				ch:     nil,
				reason: websocket.CloseError{},
			},
			want: nil,
			on:   func(f *fields) {},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
		{
			name: "set connected error",
			fields: &fields{
				state: mocks.NewDisconnectState(t),
			},
			args: &args{
				ch:     new(shadiaosocketio.Channel),
				reason: websocket.CloseError{},
			},
			want: state.ErrStateNil,
			on: func(f *fields) {
				f.state.EXPECT().SetConnected(false).Return(state.ErrStateNil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup mocks
			c := handler.NewDisconnect(tt.fields.state)

			if tt.on != nil {
				tt.on(tt.fields)
			}

			// run function
			got := c.Callback(tt.args.ch, tt.args.reason)

			// assert results
			assert.Equal(t, tt.want, got)

			if tt.assert != nil {
				tt.assert(t, tt.fields)
			}
		})
	}
}
