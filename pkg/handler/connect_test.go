package handler_test

import (
	"errors"
	"testing"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/mocks"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestConnect_Event(t *testing.T) {
	c := handler.NewConnect(nil)

	assert.Equal(t, handler.ConnectEvent, c.Event())
}

func TestConnect_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewConnect(nil)

	r.EXPECT().On(handler.ConnectEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestConnect_Occured(t *testing.T) {
	s := mocks.NewConnectState(t)
	c := handler.NewConnect(s)

	s.EXPECT().Connected().Return(false, state.ErrNotSetYet).Once()
	s.EXPECT().Connected().Return(true, nil).Once()

	assert.False(t, c.Occured())
	assert.True(t, c.Occured())
}

func TestConnect_Callback(t *testing.T) {
	type fields struct {
		state *mocks.ConnectState
	}
	type args struct {
		ch *shadiaosocketio.Channel
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
				state: mocks.NewConnectState(t),
			},
			args: &args{
				ch: new(shadiaosocketio.Channel),
			},
			want: nil,
			on: func(f *fields) {
				f.state.EXPECT().SetConnected(true).Return(nil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
		{
			name: "nil channel",
			fields: &fields{
				state: mocks.NewConnectState(t),
			},
			args: &args{
				ch: nil,
			},
			want: errors.New("nil channel"),
			on:   func(f *fields) {},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
		{
			name: "set connected error",
			fields: &fields{
				state: mocks.NewConnectState(t),
			},
			args: &args{
				ch: new(shadiaosocketio.Channel),
			},
			want: state.ErrStateNil,
			on: func(f *fields) {
				f.state.EXPECT().SetConnected(true).Return(state.ErrStateNil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup mocks
			c := handler.NewConnect(tt.fields.state)

			if tt.on != nil {
				tt.on(tt.fields)
			}

			// run function
			got := c.Callback(tt.args.ch)

			// assert results
			assert.Equal(t, tt.want, got)

			if tt.assert != nil {
				tt.assert(t, tt.fields)
			}
		})
	}
}
