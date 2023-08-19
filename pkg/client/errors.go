package client

import "errors"

// ErrTimeout is returned when a timeout occurs.
var ErrTimeout = errors.New("timeout occurred")
