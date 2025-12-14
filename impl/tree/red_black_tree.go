package tree

import (
	"github.com/harry713j/dsa_practice/constraints"
)

type Color byte

const (
	Red Color = iota
	Black
)

// Red-Black Node
type RBNode[T constraints.Ordered] struct {
	key    T
	parent *RBNode[T]
	left   *RBNode[T]
	right  *RBNode[T]
	color  Color
}

func (n *RBNode[T]) Key() T {
	return n.key
}

func (n *RBNode[T]) Left() Node[T] {
	return n.left
}

func (n *RBNode[T]) Right() Node[T] {
	return n.right
}

func (n *RBNode[T]) Parent() Node[T] {
	return n.parent
}

func (n *RBNode[T]) Color() Color {
	return n.color
}

// Red-Black tree
type RBTree[T constraints.Ordered] struct {
	Root *RBNode[T]
	NIL  *RBNode[T] // Sentinel node
}

func NewRBTree[T constraints.Ordered]() *RBTree[T] {
	nilNode := &RBNode[T]{color: Black}
	return &RBTree[T]{
		Root: nilNode,
		NIL:  nilNode,
	}
}

// Return true if the Red-Black tree is empty
func (t *RBTree[T]) Empty() bool {
	return t.Root == t.NIL
}

// Insert keys to the Red-Black tree
func (t *RBTree[T]) Insert(keys ...T) {
	for _, key := range keys {
		t.insertHelper(t.Root, key)
	}
}

// Get returns the reference to the node if the key exists in the Red-Black tree
func (t *RBTree[T]) Get(key T) (*RBNode[T], bool) {
	return t.searchHelper(key)
}

// Has returns true if the key exists in the Red-Black tree
func (t *RBTree[T]) Has(key T) bool {
	_, ok := t.searchHelper(key)

	return ok
}

// TODO: DO THE DELETE ALGORITHM AGAIN, I JUST COPY PASTED IT FROM `THE ALGORITHMS GO` WITH ONLY VAGUE UNDERSTANDING OF CONCEPT
func (t *RBTree[T]) Delete(key T) bool {
	// delete the node
	return t.deleteHelper(t.Root, key)
}

// Returns the depth of the RBTree
func (t *RBTree[T]) Depth() int {
	return calculateDepthHelper(t.Root, t.NIL)
}

// PreOrder returns the slice of value traversed in pre-order manner
func (t *RBTree[T]) PreOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return preOrderHelper(t.Root, t.NIL)
}

// InOrder returns the slice of value traversed in in-order manner
func (t *RBTree[T]) InOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return inOrderHelper(t.Root, t.NIL)
}

// PostOrder returns the slice of value traversed in post-order manner
func (t *RBTree[T]) PostOrder() []T {
	// empty tree
	if t.Empty() {
		return []T{}
	}

	return postOrderTwoStackHelper(t.Root, t.NIL)
}

// LevelOrder returns the slice of value traversed in level-order manner
func (t *RBTree[T]) LevelOrder() []T {
	if t.Empty() {
		return []T{}
	}

	return levelOrderHelper(t.Root, t.NIL)
}

func (t *RBTree[T]) insertHelper(node *RBNode[T], key T) {
	p := t.NIL // to keep track of the parent

	for node != t.NIL {
		p = node

		if key > node.key {
			node = node.right
		} else if key < node.key {
			node = node.left
		} else {
			return // node with this key already exists
		}
	}

	c := &RBNode[T]{
		key:    key,
		left:   t.NIL,
		right:  t.NIL,
		parent: p,
		color:  Red,
	}

	// if the parent is still nil
	if p == t.NIL {
		t.Root = c
	} else if c.key < p.key {
		p.left = c
	} else {
		p.right = c
	}

	// if inserted to root
	if c.parent == t.NIL {
		c.color = Black
		return
	}

	if c.parent.parent == t.NIL {
		return
	}

	// fix the node
	t.insertFix(node)
}

func (t *RBTree[T]) searchHelper(key T) (*RBNode[T], bool) {
	if t.Empty() {
		return t.NIL, false
	}

	curr := t.Root

	for curr != t.NIL {
		if key > curr.key {
			curr = curr.right
		} else if key < curr.key {
			curr = curr.left
		} else {
			return curr, true
		}
	}

	return t.NIL, false
}

func (t *RBTree[T]) insertFix(node *RBNode[T]) {

	for node.parent.color == Red {
		p := node.parent
		g := p.parent

		if p == g.left {
			u := g.right
			if u.color == Red {
				// recolor
				p.color = Black
				u.color = Black
				g.color = Red

				node = g // for above fix up
			} else {
				if node == p.right {
					node = p
					// rotate left and make it left-left
					t.leftRotate(node)
				}

				node.parent.color = Black
				node.parent.parent.color = Red
				t.rightRotate(node.parent.parent)
			}

		} else {
			u := g.left
			if u.color == Red {
				// recolor
				p.color = Black
				u.color = Black
				g.color = Red

				node = g // move above
			} else {
				if node == p.left {
					node = p
					// rotate right and make it right-right
					t.rightRotate(node)
				}

				node.parent.color = Black
				node.parent.parent.color = Red
				t.leftRotate(node.parent.parent)
			}
		}

		if node == t.Root {
			break
		}
	}

	// make root black
	t.Root.color = Black
}

func (t *RBTree[T]) leftRotate(node *RBNode[T]) {
	child := node.right
	node.right = child.left
	// change the parent
	if child.left != t.NIL {
		child.left.parent = node
	}

	child.parent = node.parent
	if node.parent == t.NIL {
		// new root
		t.Root = child
	} else if node == node.parent.left {
		node.parent.left = child
	} else {
		node.parent.right = child
	}

	child.left = node
	node.parent = child
}

func (t *RBTree[T]) rightRotate(node *RBNode[T]) {
	child := node.left
	node.left = child.right
	// change the parent
	if child.right != t.NIL {
		child.right.parent = node
	}

	child.parent = node.parent
	if node.parent == t.NIL {
		// new root
		t.Root = child
	} else if node == node.parent.left {
		node.parent.left = child
	} else {
		node.parent.right = child
	}

	child.right = node
	node.parent = child
}

// TODO: DO THE DELETE ALGORITHM AGAIN, I JUST COPY PASTED IT FROM `THE ALGORITHMS GO` WITH ONLY VAGUE UNDERSTANDING OF CONCEPT

func (t *RBTree[T]) deleteHelper(node *RBNode[T], key T) bool {
	z := t.NIL
	for node != t.NIL {
		switch {
		case node.key == key:
			z = node
			fallthrough
		case node.key <= key:
			node = node.right
		case node.key > key:
			node = node.left
		}
	}

	if z == t.NIL {
		return false
	}

	var x *RBNode[T]
	y := z
	yOriginColor := y.color
	if z.left == t.NIL {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.NIL {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = minimum[T](z.right, t.NIL).(*RBNode[T])
		yOriginColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginColor == Black {
		t.deleteFix(x)
	}

	return true
}

func (t *RBTree[T]) deleteFix(x *RBNode[T]) {
	var s *RBNode[T]
	for x != t.Root && x.color == Black {
		if x == x.parent.left {
			s = x.parent.right
			if s.color == Red {
				s.color = Black
				x.parent.color = Red
				t.leftRotate(x.parent)
				s = x.parent.right
			}

			if s.left.color == Black && s.right.color == Black {
				s.color = Red
				x = x.parent
			} else {
				if s.right.color == Black {
					s.left.color = Black
					s.color = Red
					t.rightRotate(s)
					s = x.parent.right
				}

				s.color = x.parent.color
				x.parent.color = Black
				s.right.color = Black
				t.leftRotate(x.parent)
				x = t.Root
			}
		} else {
			s = x.parent.left
			if s.color == Red {
				s.color = Black
				x.parent.color = Red
				t.rightRotate(x.parent)
				s = x.parent.left
			}

			if s.right.color == Black && s.left.color == Black {
				s.color = Red
				x = x.parent
			} else {
				if s.left.color == Black {
					s.right.color = Black
					s.color = Red
					t.leftRotate(s)
					s = x.parent.left
				}

				s.color = x.parent.color
				x.parent.color = Black
				s.left.color = Black
				t.rightRotate(x.parent)
				x = t.Root
			}
		}
	}

	x.color = Black
}

func (t *RBTree[T]) transplant(u, v *RBNode[T]) {
	switch {
	case u.parent == t.NIL:
		t.Root = v
	case u == u.parent.left:
		u.parent.left = v
	default:
		u.parent.right = v
	}

	v.parent = u.parent
}

func minimum[T constraints.Ordered](node, nilNode Node[T]) Node[T] {
	if node == nilNode {
		return node
	}

	for node.Left() != nilNode {
		node = node.Left()
	}
	return node
}
