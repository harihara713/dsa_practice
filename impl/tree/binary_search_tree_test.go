package tree

import (
	"testing"
)

func TestBSTInsertAndBasicSearch(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 3, 7, 15, 30)

	tests := []struct {
		key  int
		want bool
	}{
		{10, true},
		{5, true},
		{20, true},
		{100, false},
	}

	for _, tc := range tests {
		if ok := bst.Has(tc.key); ok != tc.want {
			t.Errorf("Has(%d) = %v, want %v", tc.key, ok, tc.want)
		}
	}
}

func TestBSTNoDuplicateInsert(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 10, 10)

	order := bst.InOrder()
	if len(order) != 1 || order[0] != 10 {
		t.Errorf("Duplicate insert failed, got inorder = %v", order)
	}
}

func TestBSTInOrderTraversal(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 3, 7, 15, 30)

	got := bst.InOrder()
	want := []int{3, 5, 7, 10, 15, 20, 30}

	if !equalSlices(got, want) {
		t.Errorf("InOrder() = %v, want %v", got, want)
	}
}

func TestBSTPreOrderTraversal(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 3, 7)

	got := bst.PreOrder()
	want := []int{10, 5, 3, 7, 20}

	if !equalSlices(got, want) {
		t.Errorf("PreOrder() = %v, want %v", got, want)
	}
}

func TestBSTPostOrderTraversal(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 3, 7)

	got := bst.PostOrder()
	want := []int{3, 7, 5, 20, 10}

	if !equalSlices(got, want) {
		t.Errorf("PostOrder() = %v, want %v", got, want)
	}
}

func TestBSTLevelOrderTraversal(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 3, 7)

	got := bst.LevelOrder()
	want := []int{10, 5, 20, 3, 7}

	if !equalSlices(got, want) {
		t.Errorf("LevelOrder() = %v, want %v", got, want)
	}
}

func TestBSTDeleteLeafNode(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 3)

	ok := bst.Delete(3)
	if !ok {
		t.Fatalf("Delete(3) failed")
	}

	got := bst.InOrder()
	want := []int{5, 10, 20}

	if !equalSlices(got, want) {
		t.Errorf("After deleting leaf, got %v, want %v", got, want)
	}
}

func TestBSTDeleteNodeWithOneChild(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 15)

	ok := bst.Delete(20)
	if !ok {
		t.Fatalf("Delete(20) failed")
	}

	got := bst.InOrder()
	want := []int{5, 10, 15}

	if !equalSlices(got, want) {
		t.Errorf("Delete one-child node, got %v, want %v", got, want)
	}
}

func TestBSTDeleteNodeWithTwoChildren(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 3, 7, 15, 30)

	ok := bst.Delete(20)
	if !ok {
		t.Fatalf("Delete(20) failed")
	}

	// After deleting 20, it should be replaced by successor(30) or predecessor(15)
	got := bst.InOrder()

	if !isInOrderBST(got) {
		t.Errorf("InOrder traversal is not sorted after deletion: %v", got)
	}
}

func TestBSTDeleteRoot(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20)

	ok := bst.Delete(10)
	if !ok {
		t.Fatalf("Delete root failed")
	}

	got := bst.InOrder()
	want := []int{5, 20}

	if !equalSlices(got, want) {
		t.Errorf("Delete root, got %v, want %v", got, want)
	}
}

func TestBSTDepth(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10, 5, 20, 3, 7)

	depth := bst.Depth()

	// Expected:
	//       10
	//      /  \
	//     5    20
	//    / \
	//   3   7
	// depth = 3
	if depth != 3 {
		t.Errorf("Depth() = %d, want 3", depth)
	}
}

func TestBSTDeleteNonExisting(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(1, 2, 3)

	ok := bst.Delete(100)
	if ok {
		t.Errorf("Delete(100) = true, want false")
	}
}

func TestEmptyBSTTraversals(t *testing.T) {
	bst := NewBST[int]()

	if len(bst.InOrder()) != 0 ||
		len(bst.PreOrder()) != 0 ||
		len(bst.PostOrder()) != 0 ||
		len(bst.LevelOrder()) != 0 {
		t.Errorf("Traversals on empty BST should return empty slices")
	}
}

//
// --- Helpers ---
//

func equalSlices[T comparable](a, b []T) bool {
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

func isInOrderBST(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
