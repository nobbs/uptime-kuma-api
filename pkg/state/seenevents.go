package state

type SeenEvents map[string]struct{}

// HasSeen returns true if the given event has occurred.
func (oe SeenEvents) HasSeen(event string) bool {
	_, ok := oe[event]
	return ok
}

// MarkSeen marks the given event as occurred.
func (oe SeenEvents) MarkSeen(event string) {
	if oe == nil {
		return
	}

	oe[event] = struct{}{}
}
