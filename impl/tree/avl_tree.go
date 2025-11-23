package tree

import (
	"github.com/harry713j/dsa_practice/constraints"
)

// Node for the AVL tree
type AVLNode[T constraints.Ordered] struct {
	key    T
	left   *AVLNode[T]
	right  *AVLNode[T]
	height int
}

func newAVLNode[T constraints.Ordered](key T, left *AVLNode[T], right *AVLNode[T]) *AVLNode[T] {
	return &AVLNode[T]{key: key, left: left, right: right, height: 1}
}

func (n *AVLNode[T]) Key() T {
	return n.key
}

func (n *AVLNode[T]) Left() Node[T] {
	return n.left
}

func (n *AVLNode[T]) Right() Node[T] {
	return n.right
}

func (n *AVLNode[T]) Height() int {
	return n.height
}

// AVL tree is the self-balancing binary tree, in AVL, the height difference between left and right side child or balance factor should
// always between {-1, 0, 1}

// AVL tree, height balancing tree
type AVL[T constraints.Ordered] struct {
	Root *AVLNode[T]
	NIL  *AVLNode[T] // a sentinel node for nil
}

// constructor func
func NewAVL[T constraints.Ordered]() *AVL[T] {
	return &AVL[T]{
		Root: nil,
		NIL:  nil,
	}
}

// Empty checks the AVL tree is empty or not
func (t *AVL[T]) Empty() bool {
	return t.Root == t.NIL
}

// when insert or delete the node in AVL tree, we need to fix it after the very operation because those operation could result a
// un-balanced tree, so its like a running a background job after each insertion and deletion

// Insert add node with provided keys
func (t *AVL[T]) Insert(keys ...T) {
	for _, key := range keys {
		t.Root = t.insertHelper(t.Root, key)
	}
}

// Get returns the node and true if that node exists in the AVL tree, otherwise return nil and false
func (t *AVL[T]) Get(key T) (*AVLNode[T], bool) {
	return t.searchHelper(key)
}

// Has returns true if the node exists in the AVL tree otherwise it returns false
func (t *AVL[T]) Has(key T) bool {
	_, ok := t.Get(key)

	return ok
}

// Delete remove a node with key == key and return true if the delete operation was successful otherwise returns false
func (t *AVL[T]) Delete(key T) bool {
	_, ok := t.Get(key)
	if !ok {
		return false
	}

	t.Root = t.deleteHelper(t.Root, key)
	return true
}

// Returns the depth of the AVL
func (t *AVL[T]) Depth() int {
	return calculateDepthHelper(t.Root, t.NIL)
}

// PreOrder returns the slice of value traversed in pre-order manner
func (t *AVL[T]) PreOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return preOrderHelper(t.Root, t.NIL)
}

// InOrder returns the slice of value traversed in in-order manner
func (t *AVL[T]) InOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return inOrderHelper(t.Root, t.NIL)
}

// PostOrder returns the slice of value traversed in post-order manner
func (t *AVL[T]) PostOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return postOrderTwoStackHelper(t.Root, t.NIL)
}

// LevelOrder returns the slice of value traversed in level-order manner
func (t *AVL[T]) LevelOrder() []T {
	if t.Empty() {
		return []T{}
	}

	return levelOrderHelper(t.Root, t.NIL)
}

// help insert the node with value key, and also balances the tree
// time: O(log n) space: O(log n), due to recursion
func (t *AVL[T]) insertHelper(node *AVLNode[T], key T) *AVLNode[T] {
	// similar to BST
	if node == t.NIL {
		// create a node and return
		return newAVLNode[T](key, t.NIL, t.NIL)
	}

	// if the key is greater than current node key then go right otherwise go left
	if key > node.Key() {
		node.right = t.insertHelper(node.right, key)
	} else if key < node.Key() {
		node.left = t.insertHelper(node.left, key)
	}

	// before returning the node set its height
	node.height = t.setHeight(node)

	// calculate the balance factor
	bf := t.balanceFactor(node)
	// if bf is > 1 or < -1 then we will perform rotation
	// if bf = 2 and bf of left node is 1 then perform Left-Left(LL) rotation
	// if bf = 2 and bf of left node is -1 then perform Left-Right(LR) rotation
	// if bf = -2 and bf of right node is -1 then perform Right-Right(RR) rotation
	// if bf = -2 and bf of right node is 1 then perform Right-Left(RL) rotation

	switch bf {
	case 2:
		switch t.balanceFactor(node.left) {
		case 1:
			return t.llRotation(node)
		case -1:
			return t.lrRotation(node)
		}

	case -2:
		switch t.balanceFactor(node.right) {
		case -1:
			return t.rrRotation(node)
		case 1:
			return t.rlRotation(node)
		}
	}

	return node
}

// search the node with key
func (t *AVL[T]) searchHelper(key T) (*AVLNode[T], bool) {
	// root is nil
	if t.Empty() {
		return t.NIL, false
	}

	curr := t.Root

	for curr != t.NIL {
		if curr.Key() == key {
			return curr, true
		}

		if curr.Key() > key {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return t.NIL, false
}

// delete the given node and balances the height of the tree
func (t *AVL[T]) deleteHelper(node *AVLNode[T], key T) *AVLNode[T] {
	// if the don't find the node
	if node == t.NIL {
		return node
	}

	// if it is a leaf node
	if node.Left() == t.NIL && node.Right() == t.NIL {
		// check if it is the root node
		if node == t.Root {
			t.Root = t.NIL
		}

		return t.NIL
	}

	// if key is greater than current node's key
	if key > node.key {
		node.right = t.deleteHelper(node.right, key)
	} else if key < node.key {
		node.left = t.deleteHelper(node.left, key)
	} else {
		// if we find the node
		// we need to which side has the greater height based on that we will delete
		// if the left side has the greater height then we will replace target node by its in-order predecessor
		// if the right side has the greater height then we will replace target node value by its in-order successor
		var leftHeight, rightHeight int

		if node.left != t.NIL {
			leftHeight = node.left.height
		}

		if node.right != t.NIL {
			rightHeight = node.right.height
		}

		if leftHeight > rightHeight {
			pred := inOrderPredecessor(node.left, t.NIL)
			// replace the key
			node.key = pred.Key()

			node.left = t.deleteHelper(node.left, pred.Key())
		} else {
			suc := inOrderSuccessor(node.right, t.NIL)

			node.key = suc.Key()

			node.right = t.deleteHelper(node.right, suc.Key())
		}
	}

	// update the height after the node is removed
	node.height = t.setHeight(node)
	// calculate the balance factor
	bf := t.balanceFactor(node)
	// if bf is > 1 or < -1 then we will perform rotation
	// if bf = 2 and bf of left node is 1 then perform Left-Left(LL) rotation
	// if bf = 2 and bf of left node is -1 then perform Left-Right(LR) rotation
	// if bf = -2 and bf of right node is -1 then perform Right-Right(RR) rotation
	// if bf = -2 and bf of right node is 1 then perform Right-Left(RL) rotation

	switch bf {
	case 2:
		switch t.balanceFactor(node.left) {
		case 1:
			return t.llRotation(node)
		case -1:
			return t.lrRotation(node)
		}

	case -2:
		switch t.balanceFactor(node.right) {
		case -1:
			return t.rrRotation(node)
		case 1:
			return t.rlRotation(node)
		}
	}

	return node
}

// set the height of a node, time: O(1) space: O(1)
func (t *AVL[T]) setHeight(node *AVLNode[T]) int {
	var leftHeight, rightHeight int

	if node != t.NIL && node.left != t.NIL {
		leftHeight = node.left.height
	}

	if node != t.NIL && node.right != t.NIL {
		rightHeight = node.right.height
	}

	// check which one is greater
	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func (t *AVL[T]) balanceFactor(node *AVLNode[T]) int {
	var leftHeight, rightHeight int

	if node != t.NIL && node.left != t.NIL {
		leftHeight = node.left.height
	}

	if node != t.NIL && node.right != t.NIL {
		rightHeight = node.right.height
	}

	return leftHeight - rightHeight
}

// Left Left rotation
func (t *AVL[T]) llRotation(node *AVLNode[T]) *AVLNode[T] {
	//      a
	//    b
	//  c
	//
	//  here b will be the parent and c will at its left as it was and a will be at right,
	// c's children will not be changed, but b's right child will be a's new left child and a's right child will be remain as it was
	a := node
	b := node.left

	a.left = b.right
	b.right = a

	// modify height of a and b i.e the node and its immediate left child
	a.height = t.setHeight(a)
	b.height = t.setHeight(b)

	// if the node is the root node, then make root point to new root
	if node == t.Root {
		t.Root = b
	}

	return b
}

// Left Right rotation
func (t *AVL[T]) lrRotation(node *AVLNode[T]) *AVLNode[T] {
	//     a
	//   b
	//    c
	//
	// here c will take place of a and a will be right child of c, b will be at c's left
	// then c's left child will be b's right child and c's right child will be a's left child, b's left and a's right will remain as it was before
	// we need to update the height of all the nodes i.e a,b and c.
	a := node
	b := node.left
	c := node.left.right

	b.right = c.left
	a.left = c.right
	c.right = a
	c.left = b

	// update the height
	c.height = t.setHeight(c)
	b.height = t.setHeight(b)
	a.height = t.setHeight(a)

	// if the node is root
	if node == t.Root {
		t.Root = c
	}

	return c
}

// Right Right rotation
func (t *AVL[T]) rrRotation(node *AVLNode[T]) *AVLNode[T] {
	//    a
	//		b
	//		  c
	//
	// here b will take the a's place and a will b's left child, b's right will be c as it was, c remains unchanged, b's left child will be a's right child
	// we need to update the height of a and b, as in LL rotation
	a := node
	b := node.right

	a.right = b.left
	b.left = a

	// update the height
	a.height = t.setHeight(a)
	b.height = t.setHeight(b)

	// if node is pointing to root
	if node == t.Root {
		t.Root = b
	}

	return b
}

// Right Left rotation
func (t *AVL[T]) rlRotation(node *AVLNode[T]) *AVLNode[T] {
	//	 a
	//		b
	//	  c
	//
	// here c will take a's place and c's left will be a, c's right b, a's right c's left and b's left will be c's right
	// we need to update the height of all three nodes i.e a, b, and c
	a := node
	b := node.right
	c := node.right.left

	a.right = c.left
	b.left = c.right

	c.left = a
	c.right = b

	// update the height
	a.height = t.setHeight(a)
	b.height = t.setHeight(b)
	c.height = t.setHeight(c)

	// if node is root
	if node == t.Root {
		t.Root = c
	}

	return c
}
