package client

import (
	"context"
	"fmt"
	"time"

	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/Baiguoshuai1/shadiaosocketio/websocket"
	"github.com/nobbs/uptime-kuma-api/pkg/handler"
	"github.com/nobbs/uptime-kuma-api/pkg/state"
)

const (
	defaultAwaitSleep = time.Duration(100) * time.Millisecond
)

// Client is the main client struct that is used to communicate with the server. It wraps the
// socket.io client, keeps track of the state and provides methods to interact with the server.
type Client struct {
	// socketio is the underlying socket.io client connection.
	socketio Connection

	// State storage for the client.
	state *state.State

	// all known handlers
	knownHandlers map[string]EventHandler
}

// Connection is the interface that wraps the basic socket.io connection methods. As socket.io is
// the only supported connection type, this interface is only used to make testing easier.
type Connection interface {
	// Ack sends an event to the server and waits for an acknowledgement. The timeout specifies how
	// long to wait for the acknowledgement before returning an error.
	Ack(string, time.Duration, ...any) (any, error)

	// On registers a handler for the given event.
	On(string, any) error

	// Close closes the connection.
	Close()
}

// EventHandler is the interface that wraps the basic handler methods required by the client.
type EventHandler interface {
	// Event returns the event name.
	Event() string

	// Register registers the handler with the connection.
	Register(handler.HandlerRegistrator) error

	// Occured returns true if the event has occured, false otherwise. Required in some places to
	// make sure a specific event has been sent before continuing.
	Occured() bool
}

// NewClient creates a new client instance and connects to the server. Returns an error if the
// connection fails.
func NewClient(host string, port int, secure bool) (c *Client, err error) {
	// create new socket.io client - this will connect to the server automatically
	var socketio *shadiaosocketio.Client
	if socketio, err = shadiaosocketio.Dial(
		shadiaosocketio.GetUrl(host, port, secure),
		*websocket.GetDefaultWebsocketTransport(),
	); err != nil {
		return nil, fmt.Errorf("socket.io client creation failed: %w", err)
	}

	// create new client instance with the socket.io connection
	return NewClientWithConnection(socketio)
}

func NewClientWithConnection(socketio Connection) (c *Client, err error) {
	// create new state instance
	s := state.NewState()

	// initialize handlers
	knownHandlers := map[string]EventHandler{
		handler.ConnectEvent:                handler.NewConnect(s),
		handler.DisconnectEvent:             handler.NewDisconnect(s),
		handler.MessageEvent:                handler.NewMessage(s),
		handler.ErrorEvent:                  handler.NewError(s),
		handler.InfoEvent:                   handler.NewInfo(s),
		handler.HeartbeatListEvent:          handler.NewHeartbeatList(s),
		handler.ImportantHeartbeatListEvent: handler.NewImportantHeartbeatList(s),
		handler.HeartbeatEvent:              handler.NewHeartbeat(s),
	}

	// create new client instance
	c = &Client{
		socketio:      socketio,
		state:         s,
		knownHandlers: knownHandlers,
	}

	// register handlers
	if err = c.registerHandlers(); err != nil {
		return nil, err
	}

	return c, nil
}

// State returns the state of the client.
func (c *Client) State() *state.State {
	return c.state
}

// Emit sends an event to the server and waits for an acknowledgement. The timeout specifies how
// long to wait for the acknowledgement before returning an error.
func (c *Client) Emit(event string, timeout time.Duration, args ...any) (any, error) {
	return c.socketio.Ack(event, timeout, args...)
}

// On registers a handler for the given event.
func (c *Client) On(event string, handler any) error {
	return c.socketio.On(event, handler)
}

// Close closes the client connection.
func (c *Client) Close() {
	c.socketio.Close()
}

// Await waits for the given event to occur for the first time. The timeout specifies how long to
// wait for the event before returning an error.
func (c *Client) Await(event string, timeout time.Duration) error {
	// Create a context with the specified timeout.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// Ensure to cancel the context to release resources.
	defer cancel()

	// Create a channel to receive a signal that work is done.
	eventChannel := make(chan bool, 1)

	// Start a goroutine to perform the check asynchronously.
	go func() {
		for {
			if ok := c.knownHandlers[event].Occured(); ok {
				eventChannel <- ok
				return
			}
			time.Sleep(defaultAwaitSleep)
		}
	}()

	select {
	case <-eventChannel:
		return nil
	case <-ctx.Done():
		return ErrTimeout
	}
}

// registerHandlers registers all known handlers with the client.
func (c *Client) registerHandlers() error {
	for _, h := range c.knownHandlers {
		if err := h.Register(c); err != nil {
			return fmt.Errorf("registering %s handler: %w", h.Event(), err)
		}
	}

	return nil
}
