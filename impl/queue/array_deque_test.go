package main

import (
	"testing"
)

func TestArrayDeque_EnqueueLast(t *testing.T) {
	d := NewArrayDeque(3)

	// Should insert elements at rear
	if err := d.EnqueueLast(10); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := d.EnqueueLast(20); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// PeekLast should be 20
	v, err := d.PeekLast()
	if err != nil || v != 20 {
		t.Fatalf("expected 20, got %d (err=%v)", v, err)
	}
}

func TestArrayDeque_EnqueueLast_Full(t *testing.T) {
	d := NewArrayDeque(2)

	_ = d.EnqueueLast(1)
	_ = d.EnqueueLast(2)

	// Third insert should fail
	if err := d.EnqueueLast(3); err == nil {
		t.Fatalf("expected full error")
	}
}

func TestArrayDeque_DequeueFirst(t *testing.T) {
	d := NewArrayDeque(3)

	_ = d.EnqueueLast(10)
	_ = d.EnqueueLast(20)

	v, err := d.DequeueFirst()
	if err != nil || v != 10 {
		t.Fatalf("expected 10, got %d (err=%v)", v, err)
	}

	v, err = d.DequeueFirst()
	if err != nil || v != 20 {
		t.Fatalf("expected 20, got %d (err=%v)", v, err)
	}

	// Now empty
	_, err = d.DequeueFirst()
	if err == nil {
		t.Fatalf("expected empty error")
	}
}

func TestArrayDeque_DequeueLast(t *testing.T) {
	d := NewArrayDeque(3)

	_ = d.EnqueueLast(10)
	_ = d.EnqueueLast(20)

	v, err := d.DequeueLast()
	if err != nil || v != 20 {
		t.Fatalf("expected 20, got %d (err=%v)", v, err)
	}

	v, err = d.DequeueLast()
	if err != nil || v != 10 {
		t.Fatalf("expected 10, got %d (err=%v)", v, err)
	}

	_, err = d.DequeueLast()
	if err == nil {
		t.Fatalf("expected empty error")
	}
}

func TestArrayDeque_Empty(t *testing.T) {
	d := NewArrayDeque(3)

	if !d.Empty() {
		t.Fatalf("expected empty")
	}

	_ = d.EnqueueLast(10)
	if d.Empty() {
		t.Fatalf("expected non-empty")
	}
}

func TestArrayDeque_Size(t *testing.T) {
	d := NewArrayDeque(5)

	if d.Size() != 0 {
		t.Fatalf("expected size 0")
	}

	_ = d.EnqueueLast(1)
	_ = d.EnqueueLast(2)

	if d.Size() != 2 {
		t.Fatalf("expected size 2, got %d", d.Size())
	}
}
