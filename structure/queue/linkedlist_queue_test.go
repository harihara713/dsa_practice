package queue

import "testing"

func TestQueue_Int(t *testing.T) {
	q := NewQueue[int]()

	// enqueue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// dequeue order
	tests := []int{1, 2, 3}
	for _, expected := range tests {
		got := q.Dequeue()
		if got != expected {
			t.Fatalf("expected %v, got %v", expected, got)
		}
	}

	if !q.Empty() {
		t.Fatal("expected queue to be empty")
	}
}

func TestQueue_String(t *testing.T) {
	q := NewQueue[string]()

	q.Enqueue("a")
	q.Enqueue("b")

	if q.Peek() != "a" {
		t.Fatalf("expected peek=a, got %v", q.Peek())
	}

	if q.Dequeue() != "a" {
		t.Fatal("expected a")
	}

	if q.Dequeue() != "b" {
		t.Fatal("expected b")
	}

	if !q.Empty() {
		t.Fatal("queue should be empty")
	}
}

func TestQueue_Bool(t *testing.T) {
	q := NewQueue[bool]()

	q.Enqueue(true)
	q.Enqueue(false)
	q.Enqueue(true)

	expected := []bool{true, false, true}
	for _, exp := range expected {
		if got := q.Dequeue(); got != exp {
			t.Fatalf("expected %v, got %v", exp, got)
		}
	}

	if !q.Empty() {
		t.Fatal("expected queue to be empty")
	}
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue[int]()

	// peek empty should give zero int → 0
	if v := q.Peek(); v != 0 {
		t.Fatalf("expected zero value, got %v", v)
	}

	q.Enqueue(10)
	q.Enqueue(20)

	if v := q.Peek(); v != 10 {
		t.Fatalf("expected peek=10, got %v", v)
	}

	// ensure size unchanged
	if q.Size() != 2 {
		t.Fatalf("expected size=2, got %d", q.Size())
	}
}

func TestQueue_Empty(t *testing.T) {
	q := NewQueue[int]()

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
	q := NewQueue[int]()

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

func TestQueue_Dequeue_Empty(t *testing.T) {
	q := NewQueue[string]()

	val := q.Dequeue() // zero value → ""
	if val != "" {
		t.Fatalf("expected zero value when dequeue empty, got %v", val)
	}
}
