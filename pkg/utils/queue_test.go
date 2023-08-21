package utils_test

import (
	"testing"

	"github.com/nobbs/uptime-kuma-api/pkg/utils"
)

func TestNewQueue(t *testing.T) {
	q := utils.NewQueue[int]()

	if q == nil {
		t.Errorf("Queue should not be nil")
	}

	if q.Len() != 0 {
		t.Errorf("Queue length should be 0, got %d", q.Len())
	}

	if !q.Empty() {
		t.Errorf("Queue should be empty")
	}

	if q.Peek() != nil {
		t.Errorf("Peek should return nil, got %d", q.Peek())
	}

	if q.Pop() != nil {
		t.Errorf("Pop should return nil, got %d", q.Pop())
	}
}

func TestNewQueueFromSlice(t *testing.T) {
	s := []int{1, 2, 3}
	q := utils.NewQueueFromSlice(s)

	if q == nil {
		t.Errorf("Queue should not be nil")
	}

	if q.Len() != 3 {
		t.Errorf("Queue length should be 3, got %d", q.Len())
	}

	if q.Empty() {
		t.Errorf("Queue should not be empty")
	}

	if *q.Peek() != 1 {
		t.Errorf("Peek should return 1, got %d", q.Peek())
	}

	if *q.Pop() != 1 {
		t.Errorf("Pop should return 1, got %d", q.Pop())
	}
}

func TestQueue_Int(t *testing.T) {
	q := utils.NewQueue[int]()
	q.Push(utils.NewInt(1))
	q.Push(utils.NewInt(2))
	q.Push(utils.NewInt(3))

	if q.Len() != 3 {
		t.Errorf("Queue length should be 3, got %d", q.Len())
	}

	if *q.Peek() != 1 {
		t.Errorf("Peek should return 1, got %d", q.Peek())
	}

	if *q.Pop() != 1 {
		t.Errorf("Pop should return 1, got %d", q.Pop())
	}

	if q.Len() != 2 {
		t.Errorf("Queue length should be 2, got %d", q.Len())
	}

	if q.Empty() {
		t.Errorf("Queue should not be empty")
	}

	q.Clear()

	if !q.Empty() {
		t.Errorf("Queue should be empty")
	}

	if q.Len() != 0 {
		t.Errorf("Queue length should be 0, got %d", q.Len())
	}

	if q.Peek() != nil {
		t.Errorf("Peek should return nil, got %d", q.Peek())
	}
}

func TestQueue_String(t *testing.T) {
	q := utils.NewQueue[string]()
	q.Push(utils.NewString("hello"))
	q.Push(utils.NewString("world"))

	if q.Len() != 2 {
		t.Errorf("Queue length should be 2, got %d", q.Len())
	}

	if *q.Peek() != "hello" {
		t.Errorf("Peek should return 'hello', got %s", *q.Peek())
	}

	if *q.Pop() != "hello" {
		t.Errorf("Pop should return 'hello', got %s", *q.Pop())
	}

	q.Push(utils.NewString("foo"))

	if q.Len() != 2 {
		t.Errorf("Queue length should be 2, got %d", q.Len())
	}

	if *q.Peek() != "world" {
		t.Errorf("Peek should return 'world', got %s", *q.Peek())
	}

	if *q.Pop() != "world" {
		t.Errorf("Pop should return 'world', got %s", *q.Pop())
	}

	if q.Len() != 1 {
		t.Errorf("Queue length should be 1, got %d", q.Len())
	}

	if *q.Peek() != "foo" {
		t.Errorf("Peek should return 'foo', got %s", *q.Peek())
	}

	if *q.Pop() != "foo" {
		t.Errorf("Pop should return 'foo', got %s", *q.Pop())
	}

	if q.Len() != 0 {
		t.Errorf("Queue length should be 0, got %d", q.Len())
	}

	if q.Peek() != nil {
		t.Errorf("Peek should return nil, got %s", *q.Peek())
	}

	if q.Pop() != nil {
		t.Errorf("Pop should return nil, got %s", *q.Pop())
	}

	if !q.Empty() {
		t.Errorf("Queue should be empty")
	}

	q.Clear()
}

func TestQueue_Trim(t *testing.T) {
	q := utils.NewQueue[string]()
	q.Push(utils.NewString("hello"))
	q.Push(utils.NewString("world"))
	q.Push(utils.NewString("foo"))
	q.Push(utils.NewString("bar"))

	if q.Len() != 4 {
		t.Errorf("Queue length should be 4, got %d", q.Len())
	}

	q.Trim(2)

	if q.Len() != 2 {
		t.Errorf("Queue length should be 2, got %d", q.Len())
	}

	if *q.Peek() != "foo" {
		t.Errorf("Peek should return 'foo', got %s", *q.Peek())
	}

	q.Push(utils.NewString("fizz"))
	q.Push(utils.NewString("buzz"))

	if q.Len() != 4 {
		t.Errorf("Queue length should be 4, got %d", q.Len())
	}

	q.Trim(5)

	if q.Len() != 4 {
		t.Errorf("Queue length should be 4, got %d", q.Len())
	}

	q.Trim(4)

	if q.Len() != 4 {
		t.Errorf("Queue length should be 4, got %d", q.Len())
	}

	q.Trim(3)

	if q.Len() != 3 {
		t.Errorf("Queue length should be 3, got %d", q.Len())
	}

	q.Trim(2)

	if q.Len() != 2 {
		t.Errorf("Queue length should be 2, got %d", q.Len())
	}

	q.Trim(1)

	if q.Len() != 1 {
		t.Errorf("Queue length should be 1, got %d", q.Len())
	}

	q.Trim(0)

	if q.Len() != 0 {
		t.Errorf("Queue length should be 0, got %d", q.Len())
	}

	q.Trim(-1)

	if q.Len() != 0 {
		t.Errorf("Queue length should be 0, got %d", q.Len())
	}
}

func TestQueue_Slice(t *testing.T) {
	s := []string{"hello", "world", "foo", "bar"}
	q := utils.NewQueueFromSlice(s)

	if q.Len() != 4 {
		t.Errorf("Queue length should be 4, got %d", q.Len())
	}

	if q.Empty() {
		t.Errorf("Queue should not be empty")
	}

	if *q.Peek() != "hello" {
		t.Errorf("Peek should return 'hello', got %s", *q.Peek())
	}

	if *q.Pop() != "hello" {
		t.Errorf("Pop should return 'hello', got %s", *q.Pop())
	}

	r := q.Slice()

	if len(r) != 3 {
		t.Errorf("Slice length should be 3, got %d", len(r))
	}

	if r[0] != "world" {
		t.Errorf("Slice[0] should be 'world', got %s", r[0])
	}

	if r[1] != "foo" {
		t.Errorf("Slice[1] should be 'foo', got %s", r[1])
	}

	if r[2] != "bar" {
		t.Errorf("Slice[2] should be 'bar', got %s", r[2])
	}

	r[0] = "fizz"

	if *q.Peek() != "world" {
		t.Errorf("Peek should return 'world', got %s", *q.Peek())
	}
}
