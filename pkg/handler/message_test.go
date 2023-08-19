package handler_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/mocks"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMessage_Event(t *testing.T) {
	c := handler.NewMessage(nil)

	assert.Equal(t, handler.MessageEvent, c.Event())
}

func TestMessage_Register(t *testing.T) {
	r := mocks.NewHandlerRegistrator(t)
	c := handler.NewMessage(nil)

	r.EXPECT().On(handler.MessageEvent, mock.MatchedBy(func(any) bool {
		return true
	})).Return(nil).Once()

	assert.NoError(t, c.Register(r))
}

func TestMessage_Occured(t *testing.T) {
	c := handler.NewMessage(nil)

	assert.False(t, c.Occured())
}

func TestMessage_Callback(t *testing.T) {
	c := handler.NewMessage(nil)

	assert.NoError(t, c.Callback(nil, nil))
}
