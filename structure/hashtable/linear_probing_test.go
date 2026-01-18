package hashtable

import (
	"reflect"
	"testing"
)

func TestInsertHTLP(t *testing.T) {
	keys := []int{11, 31, 13, 17, 46, 78, 67, 55}
	h := NewHTLP()

	for _, k := range keys {
		h.Insert(k)
	}

	expected := []int{
		0, 0, 0, 0, 0, 0, 46, 67, 0, 0, 0, 11, 31, 13, 0, 55, 0, 17, 78, 0,
	}

	if !reflect.DeepEqual(expected, h.table) {
		t.Fatalf("expected = %v, got = %v\n", expected, h.table)
	}
}

func TestGetHTLP(t *testing.T) {
	keys := []int{11, 31, 13, 17, 46, 78, 67, 55}
	h := NewHTLP()

	for _, k := range keys {
		h.Insert(k)
	}

	expected := 11
	got := h.Get(11)
	if got != expected {
		t.Fatalf("expected = %d, got = %d\n", expected, got)
	}

	expected = 67
	got = h.Get(67)
	if got != expected {
		t.Fatalf("expected = %d, got = %d\n", expected, got)
	}

	expected = -1
	got = h.Get(69)
	if got != expected {
		t.Fatalf("expected = %d, got = %d\n", expected, got)
	}
}

func TestContainsHTLP(t *testing.T) {
	keys := []int{11, 31, 13, 17, 46, 78, 67, 55}
	h := NewHTLP()

	for _, k := range keys {
		h.Insert(k)
	}

	expected := true
	got := h.Contains(11)
	if got != expected {
		t.Fatalf("expected = %t, got = %t\n", expected, got)
	}

	expected = true
	got = h.Contains(67)
	if got != expected {
		t.Fatalf("expected = %t, got = %t\n", expected, got)
	}

	expected = false
	got = h.Contains(69)
	if got != expected {
		t.Fatalf("expected = %t, got = %t\n", expected, got)
	}
}

func TestDeleteHTLP(t *testing.T) {
	keys := []int{11, 31, 13, 17, 46, 78, 67, 55}
	h := NewHTLP()

	for _, k := range keys {
		h.Insert(k)
	}

	expected := 31
	got := h.Delete(31)
	if got != expected {
		t.Fatalf("expected = %d, got = %d\n", expected, got)
	}

	expected = 67
	got = h.Delete(67)
	if got != expected {
		t.Fatalf("expected = %d, got = %d\n", expected, got)
	}

	expected = -1
	got = h.Delete(69)
	if got != expected {
		t.Fatalf("expected = %d, got = %d\n", expected, got)
	}

	expTable := []int{
		0, 0, 0, 0, 0, 0, 46, 0, 0, 0, 0, 11, 0, 13, 0, 55, 0, 17, 78, 0,
	}

	if !reflect.DeepEqual(expTable, h.table) {
		t.Fatalf("expected table = %v, got = %v\n", expTable, h.table)
	}
}
