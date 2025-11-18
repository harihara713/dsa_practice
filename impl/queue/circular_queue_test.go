package main

import (
	"testing"
)

func TestCircularQueue_EnqueueDequeue(t *testing.T) {
	q := NewCircularQueue(4) // effective capacity = 3 because of circular logic

	// enqueue 3 elements
	if err := q.Enqueue(10); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := q.Enqueue(20); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := q.Enqueue(30); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 4th insert should fail
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

func TestCircularQueue_Peek(t *testing.T) {
	q := NewCircularQueue(3)

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
	if q.Size() != 2 {
		t.Fatalf("expected size 2 after peek, got %d", q.Size())
	}
}

func TestCircularQueue_Empty(t *testing.T) {
	q := NewCircularQueue(3)

	if !q.Empty() {
		t.Fatal("expected empty queue")
	}

	q.Enqueue(11)
	if q.Empty() {
		t.Fatal("expected non-empty queue after enqueue")
	}

	q.Dequeue()
	if !q.Empty() {
		t.Fatal("expected empty queue after dequeue")
	}
}

func TestCircularQueue_Size(t *testing.T) {
	q := NewCircularQueue(4) // capacity 4, available = 3

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

func TestCircularQueue_WrapAround(t *testing.T) {
	q := NewCircularQueue(4) // effective capacity = 3

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// queue is now full
	q.Dequeue() // remove 1
	q.Dequeue() // remove 2

	// now rear should wrap around
	if err := q.Enqueue(4); err != nil {
		t.Fatalf("unexpected enqueue error: %v", err)
	}

	if err := q.Enqueue(5); err != nil {
		t.Fatalf("unexpected enqueue error: %v", err)
	}

	// queue should now contain: [3,4,5]
	val, _ := q.Dequeue()
	if val != 3 {
		t.Fatalf("expected 3, got %d", val)
	}

	val, _ = q.Dequeue()
	if val != 4 {
		t.Fatalf("expected 4, got %d", val)
	}

	val, _ = q.Dequeue()
	if val != 5 {
		t.Fatalf("expected 5, got %d", val)
	}
}
