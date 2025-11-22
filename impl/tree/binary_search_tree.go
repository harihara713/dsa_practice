package tree

import "github.com/harry713j/dsa_practice/constraints"

// BST node
type BSTNode[T constraints.Ordered] struct {
	key   T
	left  *BSTNode[T]
	right *BSTNode[T]
}

func newBSTNode[T constraints.Ordered](key T, left *BSTNode[T], right *BSTNode[T]) *BSTNode[T] {
	return &BSTNode[T]{key: key, left: left, right: right}
}

func (n *BSTNode[T]) Key() T {
	return n.key
}

func (n *BSTNode[T]) Left() Node[T] {
	return n.left
}

func (n *BSTNode[T]) Right() Node[T] {
	return n.right
}

// Binary search tree
type BSTree[T constraints.Ordered] struct {
	Root *BSTNode[T] // root of the BST
	NIL  *BSTNode[T] // sentinel nil node for comparing nil
}

// Create a new BST(Binary Search Tree)
func NewBST[T constraints.Ordered]() *BSTree[T] {
	return &BSTree[T]{Root: nil, NIL: nil}
}

// Empty checks if the BST is empty or not
func (t *BSTree[T]) Empty() bool {
	return t.Root == t.NIL
}

// Insert add a slice of node keys or value to the BST
// if the key already present then it will not add it to the BST
func (t *BSTree[T]) Insert(keys ...T) {
	for _, key := range keys {
		t.insertHelper(t.Root, key)
	}
}

// Get returns the node if exists otherwise returns NIL and false
func (t *BSTree[T]) Get(key T) (*BSTNode[T], bool) {
	return t.searchHelper(key)
}

// Has returns true if the key is exists in the BST otherwise return false
func (t *BSTree[T]) Has(key T) bool {
	_, ok := t.searchHelper(key)

	return ok
}

// Delete remove a node with value == key from the BST and returns true if the operation is successful otherwise false
func (t *BSTree[T]) Delete(key T) bool {
	_, ok := t.Get(key)
	if !ok {
		return false
	}

	t.Root = t.deleteHelper(t.Root, key)
	return true
}

// Returns the depth of the BST
func (t *BSTree[T]) Depth() int {
	return calculateDepthHelper(t.Root, t.NIL)
}

// PreOrder returns the slice of value traversed in pre-order manner
func (t *BSTree[T]) PreOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return preOrderHelper(t.Root, t.NIL)
}

// InOrder returns the slice of value traversed in in-order manner
func (t *BSTree[T]) InOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return inOrderHelper(t.Root, t.NIL)
}

// PostOrder returns the slice of value traversed in post-order manner
func (t *BSTree[T]) PostOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return postOrderTwoStackHelper(t.Root, t.NIL)
}

// LevelOrder returns the slice of value traversed in level-order manner
func (t *BSTree[T]) LevelOrder() []T {
	if t.Empty() {
		return []T{}
	}

	return levelOrderHelper(t.Root, t.NIL)
}

// help insert the value, time: O(log n), space: O(1)
func (t *BSTree[T]) insertHelper(root *BSTNode[T], key T) {
	// search the correct position for the node to add, also in order to add we need a reference to its parent node
	// for that we can use a tail pointer, also a pointer to check if the key already exists or not
	node := newBSTNode(key, t.NIL, t.NIL)
	// if the BST is empty
	if t.Empty() {
		// make root point to the new node
		t.Root = node
		return
	}

	// search for the correct position
	curr := root
	var tail *BSTNode[T] // keep the parent

	for curr != t.NIL {
		// make tail point to previous
		tail = curr
		// if the keys are equal then just return no need to insert
		if curr.Key() == key {
			return
		}

		// if the key is > the curr node value, then go right otherwise go left
		if curr.Key() > key {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	// check if the node is greater or lesser than its parent
	if node.Key() > tail.Key() {
		tail.right = node
	} else {
		tail.left = node
	}
}

// recursive insertion of node time: O(log n), space: O(log n)
func (t *BSTree[T]) insertRecursiveHelper(node *BSTNode[T], key T) *BSTNode[T] {
	if node == t.NIL {
		// create the node and return it
		return newBSTNode(key, t.NIL, t.NIL)
	}

	// if key is greater
	if key > node.Key() {
		node.right = t.insertRecursiveHelper(node.right, key)
	} else if key < node.Key() {
		node.left = t.insertRecursiveHelper(node.left, key)
	}

	return node
}

// search a node in the BST
func (t *BSTree[T]) searchHelper(key T) (*BSTNode[T], bool) {
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

// delete the node
func (t *BSTree[T]) deleteHelper(node *BSTNode[T], key T) *BSTNode[T] {
	// in order to delete a node in a BST, we also need to replace that node
	// based on the height, we will choose if we want to replace it by in-order predecessor or in-order successor
	// if the right side height is greater then we will use in-order successor otherwise in-order predecessor
	// finally if it is a leaf node then we will remove it
	if node == t.NIL {
		return node
	}

	// if it is a leaf node then we will really delete the node
	if node.Left() == t.NIL && node.Right() == t.NIL {
		// it can be a single root
		if t.Root == node {
			// make the root point to nil
			t.Root = t.NIL
		}

		return t.NIL
	}

	if node.Key() > key {
		node.left = t.deleteHelper(node.left, key)
	} else if node.Key() < key {
		node.right = t.deleteHelper(node.right, key)
	} else {
		// we found the node
		// but node can have children so based on the height of each side we will delete the node
		if calculateDepthHelper(node.left, t.NIL) > calculateDepthHelper(node.right, t.NIL) {
			// we found the in order predecessor to replace it
			pred := inOrderPredecessor(node.left, t.NIL)
			// replace the key's value of the targeted node
			node.key = pred.Key()
			// now if the in-order predecessor has children then we have to do the same with it
			node.left = t.deleteHelper(node.left, pred.Key())
		} else {
			suc := inOrderSuccessor(node.right, t.NIL)
			node.key = suc.Key()
			node.right = t.deleteHelper(node.right, suc.Key())
		}
	}

	return node
}
