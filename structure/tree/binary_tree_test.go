package tree

import (
	"reflect"
	"testing"
)

func TestBiTree_Empty(t *testing.T) {
	bt := NewBinaryTree[int]()

	if !bt.Empty() {
		t.Fatal("expected tree to be empty")
	}

	bt.Insert(10)
	if bt.Empty() {
		t.Fatal("expected non-empty tree after insert")
	}
}

func TestBiTree_Insert(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(1, 2, 3)

	level := bt.LevelOrder()
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(level, expected) {
		t.Fatalf("Insert -> LevelOrder = %v, expected %v", level, expected)
	}
}

func TestBiTree_Get_Has(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(10, 20, 30)

	node, ok := bt.Get(20)
	if !ok || node.Key() != 20 {
		t.Fatalf("expected Get(20) to find node")
	}

	if !bt.Has(30) {
		t.Fatalf("expected Has(30) to be true")
	}

	if bt.Has(99) {
		t.Fatalf("expected Has(99) to be false")
	}
}

// Delete deepest node case
func TestBiTree_Delete_Leaf(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(1, 2, 3, 4)

	// Tree:
	//     1
	//    / \
	//   2   3
	//  /
	// 4

	ok := bt.Delete(4)
	if !ok {
		t.Fatal("expected Delete(4) to return true")
	}

	level := bt.LevelOrder()
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(level, expected) {
		t.Fatalf("after delete 4 -> levelOrder = %v, expected %v", level, expected)
	}
}

// Delete node with children: requires replacement with deepest
func TestBiTree_Delete_WithChildren(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(1, 2, 3, 4, 5, 6)

	// Tree:
	//          1
	//        /   \
	//       2     3
	//      / \   /
	//     4  5  6

	// delete 2 → replaced by 6
	ok := bt.Delete(2)
	if !ok {
		t.Fatal("expected Delete(2) to succeed")
	}

	// New tree LevelOrder should be:
	//     1, 6, 3, 4, 5
	got := bt.LevelOrder()
	expected := []int{1, 6, 3, 4, 5}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("after delete 2 -> levelOrder = %v, expected %v", got, expected)
	}
}

func TestBiTree_Delete_Root(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(10, 20, 30)

	// delete root (10)
	ok := bt.Delete(10)
	if !ok {
		t.Fatal("expected Delete(10) to succeed")
	}

	// deepest-rightmost = 30 → becomes new root
	got := bt.LevelOrder()
	expected := []int{30, 20}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("after deleting root: %v, expected %v", got, expected)
	}
}

func TestBiTree_Delete_OnlyOneNode(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(42)

	ok := bt.Delete(42)
	if !ok {
		t.Fatal("expected delete to succeed on single-node tree")
	}

	if !bt.Empty() {
		t.Fatal("expected tree to be empty after deleting last node")
	}
}

func TestBiTree_Delete_NotFound(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(1, 2, 3)

	ok := bt.Delete(99)
	if ok {
		t.Fatal("expected delete(99) to be false")
	}
}
