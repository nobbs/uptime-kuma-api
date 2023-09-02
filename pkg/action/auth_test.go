package action_test

import (
	"encoding/json"
	"testing"

	"github.com/nobbs/uptime-kuma-api/mocks"
	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
	"github.com/nobbs/uptime-kuma-api/pkg/xerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	type args struct {
		c        *mocks.StatefulEmiter
		username string
		password string
		token    string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool

		on     func(*args)
		assert func(*testing.T, *args)
	}{
		{
			name: "login with username and password, return token",
			args: args{
				c:        mocks.NewStatefulEmiter(t),
				username: "username",
				password: "password",
			},
			want:    "token",
			wantErr: false,

			on: func(a *args) {
				response := map[string]any{
					"ok":    true,
					"token": "token",
				}
				rawResponse, _ := json.Marshal(response)
				emitResponse := []any{rawResponse}

				a.c.EXPECT().Await(handler.ConnectEvent, mock.AnythingOfType("time.Duration")).Return(nil)
				a.c.EXPECT().Emit(action.LoginActionName, mock.AnythingOfType("time.Duration"), mock.AnythingOfType("*action.loginRequest")).Return(emitResponse, nil)
				a.c.EXPECT().State().Return(state.NewState())
			},
			assert: func(t *testing.T, a *args) {
				a.c.AssertExpectations(t)
			},
		},
		{
			name: "login with username and password, 2fa required",
			args: args{
				c:        mocks.NewStatefulEmiter(t),
				username: "username",
				password: "password",
			},
			want:    "",
			wantErr: true,

			on: func(a *args) {
				response := map[string]any{
					"ok":            false,
					"tokenRequired": true,
				}
				rawResponse, _ := json.Marshal(response)
				emitResponse := []any{rawResponse}

				a.c.EXPECT().Await(handler.ConnectEvent, mock.AnythingOfType("time.Duration")).Return(nil)
				a.c.EXPECT().Emit(action.LoginActionName, mock.AnythingOfType("time.Duration"), mock.AnythingOfType("*action.loginRequest")).Return(emitResponse, nil)
			},
			assert: func(t *testing.T, a *args) {
				a.c.AssertExpectations(t)
			},
		},
		{
			name: "wait for autologin, return token",
			args: args{
				c: mocks.NewStatefulEmiter(t),
			},
			want:    "",
			wantErr: false,

			on: func(a *args) {
				a.c.EXPECT().Await(handler.ConnectEvent, mock.AnythingOfType("time.Duration")).Return(nil)
				a.c.EXPECT().Await(handler.AutoLoginEvent, mock.AnythingOfType("time.Duration")).Return(nil)
			},
			assert: func(t *testing.T, a *args) {
				a.c.AssertExpectations(t)
			},
		},
		{
			name: "wait for autologin, timeout",
			args: args{
				c: mocks.NewStatefulEmiter(t),
			},
			want:    "",
			wantErr: true,

			on: func(a *args) {
				a.c.EXPECT().Await(handler.ConnectEvent, mock.AnythingOfType("time.Duration")).Return(nil)
				a.c.EXPECT().Await(handler.AutoLoginEvent, mock.AnythingOfType("time.Duration")).Return(xerrors.ErrTimeout)
			},
			assert: func(t *testing.T, a *args) {
				a.c.AssertExpectations(t)
			},
		},
		{
			name: "connection event timeout",
			args: args{
				c: mocks.NewStatefulEmiter(t),
			},
			want:    "",
			wantErr: true,

			on: func(a *args) {
				a.c.EXPECT().Await(handler.ConnectEvent, mock.AnythingOfType("time.Duration")).Return(xerrors.ErrTimeout)
			},
			assert: func(t *testing.T, a *args) {
				a.c.AssertExpectations(t)
			},
		},
	}

	for ti, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.on != nil {
				tt.on(&tests[ti].args)
			}

			got, err := action.Login(tt.args.c, tt.args.username, tt.args.password, tt.args.token)
			assert.Equal(t, tt.want, got, "Should return same token")
			assert.Equal(t, tt.wantErr, err != nil, "Should return same error")

			if tt.assert != nil {
				tt.assert(t, &tests[ti].args)
			}
		})
	}
}
