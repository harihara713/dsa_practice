package tree

import (
	"reflect"
	"testing"
)

func TestRBTInsertion(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert([]int{10, 20, 30, 40, 50, 60, 70, 80}...)
	res := rbt.InOrder()
	wantInOrder := []int{10, 20, 30, 40, 50, 60, 70, 80}

	if !reflect.DeepEqual(wantInOrder, res) {
		t.Errorf("expected %v, got %v\n", wantInOrder, res)
	}
}
