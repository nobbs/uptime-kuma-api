package handler_test

import (
	"testing"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/nobbs/uptime-kuma-api/mocks"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/xerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAutoLogin_Event(t *testing.T) {
	c := handler.NewAutoLogin(nil)

	assert.Equal(t, handler.AutoLoginEvent, c.Event())
}

func TestAutoLogin_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewAutoLogin(nil)

	r.EXPECT().On(handler.AutoLoginEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestAutoLogin_Occurred(t *testing.T) {
	s := mocks.NewAutoLoginState(t)
	c := handler.NewAutoLogin(s)

	s.EXPECT().HasSeen(handler.AutoLoginEvent).Return(false).Once()
	s.EXPECT().HasSeen(handler.AutoLoginEvent).Return(true).Once()

	assert.False(t, c.Occurred())
	assert.True(t, c.Occurred())
}

func TestAutoLogin_Callback(t *testing.T) {
	type fields struct {
		state *mocks.AutoLoginState
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
				state: mocks.NewAutoLoginState(t),
			},
			args: &args{
				ch: new(shadiaosocketio.Channel),
			},
			want: nil,
			on: func(f *fields) {
				f.state.EXPECT().MarkSeen(handler.AutoLoginEvent).Return()
				f.state.EXPECT().SetAutoLogin(true).Return(nil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
		{
			name: "error",
			fields: &fields{
				state: mocks.NewAutoLoginState(t),
			},
			args: &args{
				ch: new(shadiaosocketio.Channel),
			},
			want: xerrors.ErrStateNil,
			on: func(f *fields) {
				f.state.EXPECT().MarkSeen(handler.AutoLoginEvent).Return()
				f.state.EXPECT().SetAutoLogin(true).Return(xerrors.ErrStateNil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup mocks
			c := handler.NewAutoLogin(tt.fields.state)

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
