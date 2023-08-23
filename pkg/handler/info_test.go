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

func TestInfo_Event(t *testing.T) {
	c := handler.NewInfo(nil)

	assert.Equal(t, handler.InfoEvent, c.Event())
}

func TestInfo_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewInfo(nil)

	r.EXPECT().On(handler.InfoEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestInfo_Occurred(t *testing.T) {
	s := mocks.NewInfoState(t)
	c := handler.NewInfo(s)

	s.EXPECT().Info().Return(nil, state.ErrNotSetYet).Once()
	s.EXPECT().Info().Return(&state.Info{}, nil).Once()

	assert.False(t, c.Occurred())
	assert.True(t, c.Occurred())
}

func TestInfo_Callback(t *testing.T) {
	type fields struct {
		state *mocks.InfoState
	}

	type args struct {
		ch   *shadiaosocketio.Channel
		data any
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
				state: mocks.NewInfoState(t),
			},
			args: &args{
				ch: &shadiaosocketio.Channel{},
				data: map[string]any{
					"version":              "1.22.1",
					"latestVersion":        "1.22.1",
					"primaryBaseURL":       nil,
					"serverTimezone":       "Europe/Berlin",
					"serverTimezoneOffset": "+02:00",
				},
			},
			want: nil,
			on: func(f *fields) {
				f.state.EXPECT().SetInfo(&state.Info{
					Version:              utils.NewString("1.22.1"),
					LatestVersion:        utils.NewString("1.22.1"),
					PrimaryBaseURL:       nil,
					ServerTimezone:       utils.NewString("Europe/Berlin"),
					ServerTimezoneOffset: utils.NewString("+02:00"),
				}).Return(nil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
		{
			name: "invalid data",
			fields: &fields{
				state: mocks.NewInfoState(t),
			},
			args: &args{
				ch:   &shadiaosocketio.Channel{},
				data: "invalid data",
			},
			want: utils.NewString("invalid data type"),
			on:   func(f *fields) {},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertNotCalled(t, "SetInfo", mock.Anything)
			},
		},
		{
			name: "set info failed",
			fields: &fields{
				state: mocks.NewInfoState(t),
			},
			args: &args{
				ch:   &shadiaosocketio.Channel{},
				data: map[string]any{},
			},
			want: utils.NewString(state.ErrStateNil.Error()),
			on: func(f *fields) {
				f.state.EXPECT().SetInfo(&state.Info{}).Return(state.ErrStateNil).Once()
			},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertExpectations(t)
			},
		},
		{
			name: "decode failed",
			fields: &fields{
				state: mocks.NewInfoState(t),
			},
			args: &args{
				ch: &shadiaosocketio.Channel{},
				data: map[string]any{
					"version":        &[1]int{1},
					"latestVersion":  "1.22.1",
					"primaryBaseURL": nil,
				},
			},
			want: utils.NewString("decode failed: 1 error"),
			on:   func(f *fields) {},
			assert: func(t *testing.T, f *fields) {
				f.state.AssertNotCalled(t, "SetInfo", mock.Anything)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup mocks
			c := handler.NewInfo(tt.fields.state)

			if tt.on != nil {
				tt.on(tt.fields)
			}

			// run function
			got := c.Callback(tt.args.ch, tt.args.data)

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
