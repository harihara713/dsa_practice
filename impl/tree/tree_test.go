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
