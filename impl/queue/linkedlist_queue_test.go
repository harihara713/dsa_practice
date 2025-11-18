package queue

import (
	"testing"
)

func TestQueue_EnqueueDequeue(t *testing.T) {
	var q *Queue

	tests := []struct {
		enqueue []interface{}
		dequeue []interface{}
	}{
		{
			enqueue: []interface{}{1, 2, 3},
			dequeue: []interface{}{1, 2, 3},
		},
		{
			enqueue: []interface{}{"a", "b"},
			dequeue: []interface{}{"a", "b"},
		},
		{
			enqueue: []interface{}{true, false, true},
			dequeue: []interface{}{true, false, true},
		},
	}

	for _, tt := range tests {
		q = NewQueue()

		// enqueue
		for _, v := range tt.enqueue {
			q.Enqueue(v)
		}

		// dequeue and check FIFO order
		for _, expected := range tt.dequeue {
			got := q.Dequeue()
			if got != expected {
				t.Fatalf("expected %v, got %v", expected, got)
			}
		}

		// after all dequeues, queue should be empty
		if !q.Empty() {
			t.Fatal("expected queue to be empty")
		}
	}
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue()

	// peek empty
	if q.Peek() != nil {
		t.Fatal("expected nil on empty peek")
	}

	q.Enqueue(10)
	q.Enqueue(20)

	// peek must return first element
	if v := q.Peek(); v != 10 {
		t.Fatalf("expected peek=10, got %v", v)
	}

	// size should not change
	if q.Size() != 2 {
		t.Fatalf("expected size=2, got %d", q.Size())
	}
}

func TestQueue_Empty(t *testing.T) {
	q := NewQueue()
	if !q.Empty() {
		t.Fatal("expected empty queue")
	}

	q.Enqueue(42)
	if q.Empty() {
		t.Fatal("expected non-empty queue")
	}

	q.Dequeue()
	if !q.Empty() {
		t.Fatal("expected empty queue after dequeue")
	}
}

func TestQueue_Size(t *testing.T) {
	q := NewQueue()

	if q.Size() != 0 {
		t.Fatalf("expected size 0, got %d", q.Size())
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size() != 2 {
		t.Fatalf("expected size 2, got %d", q.Size())
	}

	q.Dequeue()

	if q.Size() != 1 {
		t.Fatalf("expected size 1, got %d", q.Size())
	}
}

func TestQueue_DequeueEmpty(t *testing.T) {
	q := NewQueue()

	val := q.Dequeue()
	if val != nil {
		t.Fatalf("expected nil when dequeue empty queue, got %v", val)
	}
}
