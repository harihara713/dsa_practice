package stack

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	s := NewStack[int]()

	// push values
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Size() != 3 {
		t.Fatalf("expected size 3, got %d", s.Size())
	}

	// pop 30
	val, ok := s.Pop()
	if !ok || val != 30 {
		t.Fatalf("expected 30, got %d", val)
	}

	// pop 20
	val, ok = s.Pop()
	if !ok || val != 20 {
		t.Fatalf("expected 20, got %d", val)
	}

	// pop 10
	val, ok = s.Pop()
	if !ok || val != 10 {
		t.Fatalf("expected 10, got %d", val)
	}

	// popping when empty
	_, ok = s.Pop()
	if ok {
		t.Fatalf("expected nil, got %v", val)
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack[int]()

	s.Push(5)
	s.Push(7)

	val, ok := s.Peek()
	if !ok || val != 7 {
		t.Fatalf("expected 7, got %v", val)
	}

	// ensure size didn’t change
	if s.Size() != 2 {
		t.Fatalf("expected size 2 after peek, got %d", s.Size())
	}
}

func TestStackEmpty(t *testing.T) {
	s := NewStack[int]()

	if !s.Empty() {
		t.Fatal("stack should be empty")
	}

	s.Push(10)

	if s.Empty() {
		t.Fatal("stack should NOT be empty after push")
	}
}

func TestStackSize(t *testing.T) {
	s := NewStack[int]()

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
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	expected := "3 -> 2 -> 1 -> "
	if s.String() != expected {
		t.Fatalf("expected %q, got %q", expected, s.String())
	}
}
