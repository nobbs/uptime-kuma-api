package handler_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/mocks"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestError_Event(t *testing.T) {
	c := handler.NewError(nil)

	assert.Equal(t, handler.ErrorEvent, c.Event())
}

func TestError_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewError(nil)

	r.EXPECT().On(handler.ErrorEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestError_Occurred(t *testing.T) {
	c := handler.NewError(nil)

	assert.False(t, c.Occurred())
}

func TestError_Callback(t *testing.T) {
	c := handler.NewError(nil)

	assert.NoError(t, c.Callback(nil, nil))
}
