package queue

import (
	"testing"
)

func TestStackQueue_EnqueueDequeue(t *testing.T) {
	q := NewStackQueue()

	// enqueue
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if q.Size() != 3 {
		t.Errorf("expected size 3, got %d", q.Size())
	}

	// dequeue
	v, err := q.Dequeue()
	if err != nil || v != 10 {
		t.Errorf("expected 10, got %v, err=%v", v, err)
	}

	v, err = q.Dequeue()
	if err != nil || v != 20 {
		t.Errorf("expected 20, got %v, err=%v", v, err)
	}

	v, err = q.Dequeue()
	if err != nil || v != 30 {
		t.Errorf("expected 30, got %v, err=%v", v, err)
	}

	// dequeue on empty
	_, err = q.Dequeue()
	if err == nil {
		t.Error("expected error on empty queue, got nil")
	}
}

func TestStackQueue_Peek(t *testing.T) {
	q := NewStackQueue()

	q.Enqueue(5)
	q.Enqueue(6)

	peek, err := q.Peek()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if peek != 5 {
		t.Errorf("expected peek 5, got %v", peek)
	}

	q.Dequeue()

	peek, err = q.Peek()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if peek != 6 {
		t.Errorf("expected peek 6, got %v", peek)
	}
}

func TestStackQueue_Empty(t *testing.T) {
	q := NewStackQueue()

	if !q.Empty() {
		t.Errorf("queue should be empty")
	}

	q.Enqueue(1)

	if q.Empty() {
		t.Errorf("queue should not be empty after enqueue")
	}

	q.Dequeue()

	if !q.Empty() {
		t.Errorf("queue should be empty after removing all items")
	}
}

func TestStackQueue_Size(t *testing.T) {
	q := NewStackQueue()

	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}

	q.Enqueue(10)
	q.Enqueue(20)

	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	q.Dequeue()

	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}
}
