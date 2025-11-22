package tree

import (
	"reflect"
	"testing"
)

// helper function to build a test tree:
//
//	     1
//	   /   \
//	  2     3
//	 / \   /
//	4  5  6
func buildTestTree() *BiTree[int] {
	t := NewBinaryTree[int]()
	t.Insert(1, 2, 3, 4, 5, 6)
	return t
}

func TestPreOrder(t *testing.T) {
	bt := buildTestTree()
	got := bt.PreOrder()
	expected := []int{1, 2, 4, 5, 3, 6}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("PreOrder = %v, expected %v", got, expected)
	}
}

func TestInOrder(t *testing.T) {
	bt := buildTestTree()
	got := bt.InOrder()
	expected := []int{4, 2, 5, 1, 6, 3}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("InOrder = %v, expected %v", got, expected)
	}
}

func TestPostOrder(t *testing.T) {
	bt := buildTestTree()
	got := bt.PostOrder()
	expected := []int{4, 5, 2, 6, 3, 1}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("PostOrder = %v, expected %v", got, expected)
	}
}

func TestLevelOrder(t *testing.T) {
	bt := buildTestTree()
	got := bt.LevelOrder()
	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("LevelOrder = %v, expected %v", got, expected)
	}
}

func TestDepth(t *testing.T) {
	bt := buildTestTree()
	got := bt.Depth()

	// depth = 3 levels
	if got != 3 {
		t.Fatalf("Depth = %d, expected %d", got, 3)
	}
}

func TestInOrderPredecessor(t *testing.T) {
	nilNode := (*BSTNode[int])(nil)

	bst := NewBST[int]()
	bst.Insert([]int{20, 10, 30, 25, 5, 15, 7}...)

	// predecessor of 20 should be 10
	pred := inOrderPredecessor(bst.Root.left, nilNode)
	if pred.Key() != 15 {
		t.Fatalf("expected predecessor of root to be 15, got %v", pred.Key())
	}

	// predecessor of 25 (no right subtree) is 25 itself (your logic)
	pred = inOrderPredecessor[int](bst.Root.left.left, nilNode)
	if pred.Key() != 7 {
		t.Fatalf("expected predecessor of 10 to be 7, got %v", pred.Key())
	}
}

func TestInOrderSuccessor(t *testing.T) {
	nilNode := (*BSTNode[int])(nil)

	/*
	      20
	    /    \
	   10    30
	        /  \
	       25  40
	*/

	bst := NewBST[int]()
	bst.Insert([]int{20, 10, 30, 25, 40, 35}...)

	suc := inOrderSuccessor(bst.Root.right, nilNode)
	if suc.Key() != 25 {
		t.Fatalf("expected successor of root to be 25 got %v", suc.Key())
	}

	suc = inOrderSuccessor(bst.Root.right.right, nilNode)
	if suc.Key() != 35 {
		t.Fatalf("expected successor of 30 to be 35 got %v", suc.Key())
	}
}
