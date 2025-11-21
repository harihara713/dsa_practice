package tree

import "github.com/harry713j/dsa_practice/impl/queue"

// Binary tree node
type BiTreeNode[T comparable] struct {
	key   T
	left  *BiTreeNode[T]
	right *BiTreeNode[T]
}

func newBiTreeNode[T comparable](key T, left *BiTreeNode[T], right *BiTreeNode[T]) *BiTreeNode[T] {
	return &BiTreeNode[T]{
		key:   key,
		left:  left,
		right: right,
	}
}

func (n *BiTreeNode[T]) Key() T {
	return n.key
}

func (n *BiTreeNode[T]) Left() Node[T] {
	return n.left
}

func (n *BiTreeNode[T]) Right() Node[T] {
	return n.right
}

// Binary tree
type BiTree[T comparable] struct {
	root *BiTreeNode[T]
	NIL  *BiTreeNode[T] // a sentinel value to compare for nil
}

func NewBinaryTree[T comparable]() *BiTree[T] {
	return &BiTree[T]{root: nil, NIL: nil}
}

// Empty check if the Binary Tree is empty or not
func (t *BiTree[T]) Empty() bool {
	return t.root == t.NIL
}

// Insert a node to the tree, if value == NIL of then stop inserting for that node
func (t *BiTree[T]) Insert(keys ...T) {
	for _, key := range keys {
		t.insertHelper(t.root, key)
	}
}

// Get a node from the binary tree
func (t *BiTree[T]) Get(key T) (Node[T], bool) {
	return t.searchBinaryTreeHelper(t.root, key)
}

// Has check if the key exists in the binary tree or not
func (t *BiTree[T]) Has(key T) bool {
	_, ok := t.searchBinaryTreeHelper(t.root, key)
	return ok
}

// Delete remove the node from the binary tree, if success then return true else return false
func (t *BiTree[T]) Delete(key T) bool {
	// if the root is NIL
	if t.root == t.NIL {
		return false
	}

	// if only one node
	if t.root.left == t.NIL && t.root.right == t.NIL {
		// check if root value equal to key
		if t.root.key == key {
			t.root = t.NIL
			return true
		}

		return false
	}

	// using the queue do BFS or level order traverse
	q := queue.NewQueue[*BiTreeNode[T]]()
	q.Enqueue(t.root)

	// find the node to be removed, replace that with right-most deepest node
	// keep track of the target node, deepest node and parent of that right most deepest node
	// replace the value of target with the deepest's value and remove the link from the parent of the deepest with deepest
	var target, deepest, parent *BiTreeNode[T]

	for !q.Empty() {
		p := q.Dequeue()

		// check if the value is equal or not, if equal we find our target
		if p.Key() == key {
			target = p
		}

		deepest = p // last visited node

		if p.left != t.NIL {
			parent = p
			q.Enqueue(p.left)
		}

		if p.right != t.NIL {
			parent = p
			q.Enqueue(p.right)
		}
	}

	// if the target is nil
	if target == nil {
		return false
	}

	// replace the value
	target.key = deepest.key

	if parent.left == deepest {
		parent.left = t.NIL
	} else if parent.right == deepest {
		parent.right = t.NIL
	}

	return true
}

// PreOrder returns the slice of value traversed in pre-order manner
func (t *BiTree[T]) PreOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return preOrderHelper(t.root, t.NIL)
}

// InOrder returns the slice of value traversed in in-order manner
func (t *BiTree[T]) InOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return inOrderHelper(t.root, t.NIL)
}

// PostOrder returns the slice of value traversed in post-order manner
func (t *BiTree[T]) PostOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return postOrderTwoStackHelper(t.root, t.NIL)
}

// LevelOrder returns the slice of value traversed in level-order manner
func (t *BiTree[T]) LevelOrder() []T {
	if t.Empty() {
		return []T{}
	}

	return levelOrderHelper(t.root, t.NIL)
}

// Depth of the binary tree
func (t *BiTree[T]) Depth() int {
	return calculateDepthHelper(t.root, t.NIL)
}

// helper function for inserting the key to the Binary tree
func (t *BiTree[T]) insertHelper(root *BiTreeNode[T], key T) {
	// check if root is null or not, if so then make root point to that node
	node := newBiTreeNode(key, t.NIL, t.NIL)

	if root == t.NIL {
		t.root = node
		return
	}

	// if the root is not null, using a queue of node we can check the
	q := queue.NewQueue[*BiTreeNode[T]]()
	q.Enqueue(root)

	for !q.Empty() {
		// deque the value and check if the left and right are nil or not
		// if nil then insert at that position otherwise add that node to the queue
		p := q.Dequeue()

		// left
		if p.left == t.NIL {
			p.left = node
			return
		} else {
			q.Enqueue(p.left)
		}

		// right
		if p.right == t.NIL {
			p.right = node
			return
		} else {
			q.Enqueue(p.right)
		}
	}
}

// search a node in the binary tree
func (t *BiTree[T]) searchBinaryTreeHelper(root *BiTreeNode[T], key T) (*BiTreeNode[T], bool) {
	if root == t.NIL {
		return root, false
	}

	if root.Key() == key {
		return root, true
	}

	q := queue.NewQueue[*BiTreeNode[T]]()
	q.Enqueue(root)

	for !q.Empty() {
		// deque the value
		p := q.Dequeue()
		// check the left is equal to the key, if so then return that node otherwise add it to the queue
		if p.left != t.NIL && p.left.Key() == key {
			return p.left, true
		} else if p.left != t.NIL {
			q.Enqueue(p.left)
		}

		// check the right is equal to the key, if so return the node otherwise add it to the queue
		if p.right != t.NIL && p.right.Key() == key {
			return p.right, true
		} else if p.right != t.NIL {
			q.Enqueue(p.right)
		}
	}

	// we don't find the node with key == key
	return t.NIL, false
}
