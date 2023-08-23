package utils

// Queue is a generic queue implementation based on slices.
type Queue[T any] struct {
	// Stores the queue.
	queue []T
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		queue: make([]T, 0),
	}
}

// NewQueueFromSlice creates a new queue from a slice.
func NewQueueFromSlice[T any](slice []T) *Queue[T] {
	return &Queue[T]{
		queue: slice,
	}
}

// Push adds an item to the queue.
func (q *Queue[T]) Push(item *T) {
	q.queue = append(q.queue, *item)
}

// Pop removes an item from the queue.
func (q *Queue[T]) Pop() *T {
	if q.Empty() {
		return nil
	}

	item := q.queue[0]
	q.queue = q.queue[1:]

	return &item
}

// Len returns the length of the queue.
func (q *Queue[T]) Len() int {
	return len(q.queue)
}

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

// Peek returns the first item in the queue without removing it.
func (q *Queue[T]) Peek() *T {
	if q.Empty() {
		return nil
	}

	return &q.queue[0]
}

// Clear removes all items from the queue.
func (q *Queue[T]) Clear() {
	q.queue = make([]T, 0)
}

// Trim removes all items from the queue except the last n items.
func (q *Queue[T]) Trim(n int) {
	if n > q.Len() || n < 0 {
		return
	}

	q.queue = q.queue[len(q.queue)-n:]
}

// Slice returns copy of queue data as slice.
func (q *Queue[T]) Slice() []T {
	slice := make([]T, len(q.queue))
	copy(slice, q.queue)

	return slice
}
