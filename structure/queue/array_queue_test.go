package queue

import (
	"testing"
)

func TestArrayQueue_EnqueueDequeue(t *testing.T) {
	q := NewArrayQueue(3)

	// enqueue values
	if err := q.Enqueue(10); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := q.Enqueue(20); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := q.Enqueue(30); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// queue should now be full
	if err := q.Enqueue(40); err == nil {
		t.Fatalf("expected queue full error, got nil")
	}

	// dequeue 10
	val, err := q.Dequeue()
	if err != nil || val != 10 {
		t.Fatalf("expected 10, got %d err=%v", val, err)
	}

	// dequeue 20
	val, err = q.Dequeue()
	if err != nil || val != 20 {
		t.Fatalf("expected 20, got %d err=%v", val, err)
	}

	// dequeue 30
	val, err = q.Dequeue()
	if err != nil || val != 30 {
		t.Fatalf("expected 30, got %d err=%v", val, err)
	}

	// should now be empty
	_, err = q.Dequeue()
	if err == nil {
		t.Fatalf("expected queue empty error, got nil")
	}
}

func TestArrayQueue_Peek(t *testing.T) {
	q := NewArrayQueue(5)

	// empty peek
	_, err := q.Peek()
	if err == nil {
		t.Fatalf("expected error on empty queue peek, got nil")
	}

	q.Enqueue(5)
	q.Enqueue(7)

	val, err := q.Peek()
	if err != nil || val != 5 {
		t.Fatalf("expected peek=5, got %d err=%v", val, err)
	}

	// ensure peek does not remove
	size := q.rear - q.front
	if size != 2 {
		t.Fatalf("expected size 2, got %d", size)
	}
}

func TestArrayQueue_Empty(t *testing.T) {
	q := NewArrayQueue(2)

	if !q.Empty() {
		t.Fatal("expected queue to be empty")
	}

	q.Enqueue(1)

	if q.Empty() {
		t.Fatal("expected queue to be non-empty")
	}

	q.Dequeue()

	if !q.Empty() {
		t.Fatal("expected queue to be empty after dequeue")
	}
}

func TestArrayQueue_Size(t *testing.T) {
	q := NewArrayQueue(3)

	if q.Size() != 0 {
		t.Fatalf("expected size 0, got %d", q.Size())
	}

	q.Enqueue(11)
	q.Enqueue(22)

	if q.Size() != 2 {
		t.Fatalf("expected size 2, got %d", q.Size())
	}
}
