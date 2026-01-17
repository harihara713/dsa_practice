package tree

import (
	"fmt"
	"testing"
)

//
// ---------- Helpers ----------
//

func newAVL() *AVL[int] {
	return NewAVL[int]()
}

//
// ---------- Insert Tests ----------
//

// LL Rotation Test
func TestAVLInsertLL(t *testing.T) {
	avl := newAVL()

	avl.Insert(30, 20, 10)

	got := avl.PreOrder()
	want := []int{20, 10, 30}

	if !equal(got, want) {
		t.Fatalf("LL rotation failed. got=%v want=%v", got, want)
	}
}

// RR Rotation Test
func TestAVLInsertRR(t *testing.T) {
	avl := newAVL()

	avl.Insert(10, 20, 30)

	got := avl.PreOrder()
	want := []int{20, 10, 30}

	if !equal(got, want) {
		t.Fatalf("RR rotation failed. got=%v want=%v", got, want)
	}
}

// LR Rotation Test
func TestAVLInsertLR(t *testing.T) {
	avl := newAVL()

	avl.Insert(30, 10, 20) // LR case

	got := avl.PreOrder()
	want := []int{20, 10, 30}

	if !equal(got, want) {
		t.Fatalf("LR rotation failed. got=%v want=%v", got, want)
	}
}

// RL Rotation Test
func TestAVLInsertRL(t *testing.T) {
	avl := newAVL()

	avl.Insert(10, 30, 20) // RL case

	got := avl.PreOrder()
	want := []int{20, 10, 30}

	if !equal(got, want) {
		t.Fatalf("RL rotation failed. got=%v want=%v", got, want)
	}
}

//
// ---------- Search Tests ----------
//

func TestAVLSearch(t *testing.T) {
	avl := newAVL()
	avl.Insert(10, 20, 30)

	exists := avl.Has(20)
	if !exists {
		t.Fatalf("Has failed: expected to find key 20")
	}

	exists = avl.Has(99)
	if exists {
		t.Fatalf("Has should have failed but return the true")
	}
}

//
// ---------- Delete Tests ----------
//

// delete leaf node
func TestAVLDeleteLeaf(t *testing.T) {
	avl := newAVL()
	avl.Insert(20, 10, 30)

	deleted := avl.Delete(10)

	fmt.Printf("delete result: %v\n", deleted)

	got := avl.InOrder()
	want := []int{20, 30}

	if !equal(got, want) {
		t.Fatalf("Delete leaf failed. got=%v want=%v", got, want)
	}
}

// delete node with one child
func TestAVLDeleteOneChild(t *testing.T) {
	avl := newAVL()
	avl.Insert(20, 10, 30, 25)

	avl.Delete(30)

	got := avl.InOrder()
	want := []int{10, 20, 25}

	if !equal(got, want) {
		t.Fatalf("Delete one-child node failed. got=%v want=%v", got, want)
	}
}

// delete node with two children + ensure rebalancing
func TestAVLDeleteTwoChildren(t *testing.T) {
	avl := newAVL()
	avl.Insert(50, 30, 70, 20, 40, 60, 80)

	avl.Delete(70) // node with two children

	got := avl.InOrder()
	want := []int{20, 30, 40, 50, 60, 80}

	if !equal(got, want) {
		t.Fatalf("Delete two-child node failed. got=%v want=%v", got, want)
	}

	if avl.Depth() > 3 {
		t.Fatalf("not balanced after deletion, depth=%d", avl.Depth())
	}
}

//
// ---------- Traversal Tests ----------
//

func TestAVLTraversals(t *testing.T) {
	avl := newAVL()
	avl.Insert(40, 20, 60, 10, 30, 50, 70)

	if !equal(avl.InOrder(), []int{10, 20, 30, 40, 50, 60, 70}) {
		t.Fatalf("InOrder traversal incorrect: %v", avl.InOrder())
	}

	if len(avl.PreOrder()) == 0 {
		t.Fatalf("PreOrder should not be empty")
	}

	if len(avl.PostOrder()) == 0 {
		t.Fatalf("PostOrder should not be empty")
	}

	if len(avl.LevelOrder()) == 0 {
		t.Fatalf("LevelOrder should not be empty")
	}
}

//
// ---------- Depth Tests ----------
//

func TestAVLDepth(t *testing.T) {
	avl := newAVL()
	for i := 1; i <= 100; i++ {
		avl.Insert(i)
	}

	if avl.Depth() > 10 {
		t.Fatalf("AVL depth too large; likely unbalanced: depth=%d", avl.Depth())
	}
}

//
// ---------- Duplicate Insert Test ----------
//

func TestAVLDuplicateInsert(t *testing.T) {
	avl := newAVL()
	avl.Insert(10, 10, 10)

	got := avl.InOrder()
	want := []int{10}

	if !equal(got, want) {
		t.Fatalf("Duplicate insert should be ignored. got=%v want=%v", got, want)
	}
}

//
// ---------- Empty Behavior ----------
//

func TestAVLEmpty(t *testing.T) {
	avl := newAVL()

	if !avl.Empty() {
		t.Fatalf("New should be empty")
	}

	avl.Insert(1)
	if avl.Empty() {
		t.Fatalf("should not be empty after insert")
	}
}

//
// ---------- Utility ----------
//

func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
