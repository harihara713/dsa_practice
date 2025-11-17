package main

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {
	s := NewArrayStack(3)

	// push
	if err := s.Push(10); err != nil {
		t.Fatalf("push failed: %v", err)
	}
	if err := s.Push(20); err != nil {
		t.Fatalf("push failed: %v", err)
	}
	if err := s.Push(30); err != nil {
		t.Fatalf("push failed: %v", err)
	}

	// pushing when full
	if err := s.Push(40); err != ErrStackOverflow {
		t.Fatalf("expected stack overflow error, got %v", err)
	}

	// pop
	val, err := s.Pop()
	if err != nil || val != 30 {
		t.Fatalf("expected 30, got %d (%v)", val, err)
	}

	val, err = s.Pop()
	if err != nil || val != 20 {
		t.Fatalf("expected 20, got %d (%v)", val, err)
	}

	val, err = s.Pop()
	if err != nil || val != 10 {
		t.Fatalf("expected 10, got %d (%v)", val, err)
	}

	// popping when empty
	_, err = s.Pop()
	if err != ErrStackEmpty {
		t.Fatalf("expected empty error, got %v", err)
	}
}

func TestPeek(t *testing.T) {
	s := NewArrayStack(2)

	_, err := s.Peek()
	if err != ErrStackEmpty {
		t.Fatalf("expected empty error, got %v", err)
	}

	s.Push(5)
	s.Push(7)

	val, err := s.Peek()
	if err != nil || val != 7 {
		t.Fatalf("expected 7, got %d (%v)", val, err)
	}
}

func TestEmptyAndFull(t *testing.T) {
	s := NewArrayStack(1)

	if !s.Empty() {
		t.Fatal("stack should be empty initially")
	}

	s.Push(10)
	if s.Empty() {
		t.Fatal("stack should not be empty after push")
	}

	if !s.Full() {
		t.Fatal("stack should be full with capacity 1")
	}
}

func TestSize(t *testing.T) {
	s := NewArrayStack(5)

	if s.Size() != 0 {
		t.Fatalf("expected size 0, got %d", s.Size())
	}

	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Fatalf("expected size 2, got %d", s.Size())
	}
}

func TestString(t *testing.T) {
	s := NewArrayStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	expected := "3 -> 2 -> 1 "
	if s.String() != expected {
		t.Fatalf("expected %q, got %q", expected, s.String())
	}
}
