package stack

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	s := NewStack()

	// push values
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Size() != 3 {
		t.Fatalf("expected size 3, got %d", s.Size())
	}

	// pop 30
	val := s.Pop()
	if val == nil || val != 30 {
		t.Fatalf("expected 30, got %d", val)
	}

	// pop 20
	val = s.Pop()
	if val == nil || val != 20 {
		t.Fatalf("expected 20, got %d", val)
	}

	// pop 10
	val = s.Pop()
	if val == nil || val != 10 {
		t.Fatalf("expected 10, got %d", val)
	}

	// popping when empty
	val = s.Pop()
	if val != nil {
		t.Fatalf("expected nil, got %v", val)
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()

	s.Push(5)
	s.Push(7)

	val := s.Peek()
	if val == nil {
		t.Fatalf("expected 7, got %v", val)
	}

	if val != 7 {
		t.Fatalf("expected 7, got %v", val)
	}

	// ensure size didn’t change
	if s.Size() != 2 {
		t.Fatalf("expected size 2 after peek, got %d", s.Size())
	}
}

func TestStackEmpty(t *testing.T) {
	s := NewStack()

	if !s.Empty() {
		t.Fatal("stack should be empty")
	}

	s.Push(10)

	if s.Empty() {
		t.Fatal("stack should NOT be empty after push")
	}
}

func TestStackSize(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Fatalf("expected size 0, got %d", s.Size())
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Fatalf("expected size 3, got %d", s.Size())
	}
}

func TestStackString(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	expected := "3 -> 2 -> 1 -> "
	if s.String() != expected {
		t.Fatalf("expected %q, got %q", expected, s.String())
	}
}
