package hashtable

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHashtable_BasicOperations(t *testing.T) {
	ht := New()

	tests := []struct {
		name     string
		key      any
		value    any
		wantGet  any
		wantCont bool
	}{
		{"put string key", "name", "Alice", "Alice", true},
		{"put int key", 42, "answer", "answer", true},
		{"put float key", 3.14, "pi", "pi", true},
		{"overwrite existing", "name", "Bob", "Bob", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht.Put(tt.key, tt.value)
			fmt.Printf("ht: %+v\n", ht)

			got := ht.Get(tt.key)
			if !reflect.DeepEqual(got, tt.wantGet) {
				t.Errorf("Get(%v) = %v, want %v", tt.key, got, tt.wantGet)
			}

			if ht.Contains(tt.key) != tt.wantCont {
				t.Errorf("Contains(%v) = %v, want %v", tt.key, ht.Contains(tt.key), tt.wantCont)
			}
		})
	}
}

func TestHashtable_Delete(t *testing.T) {
	ht := New()

	ht.Put("a", 1)
	ht.Put("b", 2)
	ht.Put("c", 3)

	t.Run("delete existing", func(t *testing.T) {
		got := ht.Delete("b")
		if got != 2 {
			t.Errorf("Delete('b') returned %v, want 2", got)
		}
		if ht.Contains("b") {
			t.Error("key 'b' still exists after delete")
		}
		if ht.Get("b") != nil {
			t.Error("Get('b') after delete should return nil")
		}
	})

	t.Run("delete non-existing", func(t *testing.T) {
		got := ht.Delete("xyz")
		if got != nil {
			t.Errorf("Delete non-existing key returned %v, want nil", got)
		}
	})

	t.Run("delete and check others remain", func(t *testing.T) {
		if ht.Get("a") != 1 {
			t.Error("key 'a' was affected by delete")
		}
		if ht.Get("c") != 3 {
			t.Error("key 'c' was affected by delete")
		}
	})
}

func TestHashtable_CollisionAndProbing(t *testing.T) {
	// We force collisions by using keys that hash to same index
	// (in practice this depends on the current hash function)

	ht := New()

	// These keys are likely to collide depending on fnv hash
	keys := []any{
		"cat",
		"act", // very likely to collide with "cat" in many hash functions
		"tac",
		42,
	}

	for i, k := range keys {
		ht.Put(k, i+100)
	}

	for i, k := range keys {
		t.Run(fmt.Sprintf("get after collision key=%v", k), func(t *testing.T) {
			if got := ht.Get(k); got != i+100 {
				t.Errorf("Get(%v) = %v, want %d", k, got, i+100)
			}
			if !ht.Contains(k) {
				t.Errorf("Contains(%v) = false, want true", k)
			}
		})
	}
}

func TestHashtable_Resize(t *testing.T) {
	ht := New()

	// Insert enough elements to trigger resize (load factor > 0.5)
	// defaultCapacity = 13 → resize at ~7 elements
	for i := 0; i < 10; i++ {
		ht.Put(i, fmt.Sprintf("val-%d", i))
	}

	if ht.capacity <= 13 {
		t.Errorf("capacity should have increased after resize, got %d", ht.capacity)
	}

	if ht.size != 10 {
		t.Errorf("size should be 10 after inserts, got %d", ht.size)
	}

	// Verify all keys are still accessible
	for i := 0; i < 10; i++ {
		if got := ht.Get(i); got != fmt.Sprintf("val-%d", i) {
			t.Errorf("after resize Get(%d) = %v, want val-%d", i, got, i)
		}
	}
}

func TestHashtable_DeleteAfterResize(t *testing.T) {
	ht := New()

	// Fill past resize point
	for i := 0; i < 12; i++ {
		ht.Put(i, i*10)
	}

	originalCapacity := ht.capacity

	// Delete some elements
	for i := 0; i < 5; i++ {
		ht.Delete(i)
	}

	// Capacity should NOT decrease in your current implementation
	if ht.capacity != originalCapacity {
		t.Errorf("expected capacity to remain %d after deletes, got %d", originalCapacity, ht.capacity)
	}

	// But values should be gone
	if ht.Get(2) != nil {
		t.Error("deleted key 2 still accessible")
	}
}

// func TestHashtable_ZeroCapacityEdgeCase(t *testing.T) {
// 	// Just to show behavior — normally not useful
// 	ht := &Hashtable{
// 		capacity: 0,
// 		table:    make([]*Entry, 0),
// 	}

// 	ht.Put("key", "value") // should panic or behave badly — your code doesn't handle capacity=0

// 	// Your current code will panic on modulo 0 in hash()
// 	// This test just documents current behavior (panic expected)
// 	defer func() {
// 		if r := recover(); r == nil {
// 			t.Error("expected panic on hash with capacity=0, but didn't panic")
// 		}
// 	}()

// 	_ = ht.hash("key")
// }
