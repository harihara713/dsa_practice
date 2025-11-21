package tree

import (
	"github.com/harry713j/dsa_practice/impl/queue"
)

type Node[T comparable] interface {
	Key() T
	Left() Node[T]
	Right() Node[T]
}

// preorder traversal

// Recursive PreOrder, time: O(n), space: O(log n) to O(n)
func preOrderRecursiveHelper[T comparable](node, nilNode Node[T], res []T) {
	if node == nilNode {
		return
	}

	// add the element to the res and then call left and right
	res = append(res, node.Key())
	preOrderRecursiveHelper(node.Left(), nilNode, res)
	preOrderRecursiveHelper(node.Right(), nilNode, res)
}

// Iterative Preorder
func preOrderHelper[T comparable](node, nilNode Node[T]) []T {
	// stack to store the nodes
	var stack []Node[T]
	var res []T // store the result

	for node != nilNode || len(stack) > 0 {
		// if the node is not nil then add it to result and push it to the stack and go to left, otherwise pop and go to right
		if node != nilNode {
			res = append(res, node.Key())
			stack = append(stack, node)
			node = node.Left()
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = temp.Right()
		}
	}

	return res
}

// in-order traversal

// Recursive InOrder, time: O(n) space: O(log n) to O(n)
func inOrderRecursiveHelper[T comparable](node, nilNode Node[T], res []T) {
	if node == nilNode {
		return
	}

	inOrderRecursiveHelper(node.Left(), nilNode, res)
	res = append(res, node.Key())
	inOrderRecursiveHelper(node.Right(), nilNode, res)
}

// Iterative InOrder
func inOrderHelper[T comparable](node, nilNode Node[T]) []T {
	// stack to store the nodes
	var stack []Node[T]
	var res []T // store the result

	for node != nilNode || len(stack) > 0 {
		// if the node is not nil then add it to result and push it to the stack and go to left, otherwise pop and go to right
		if node != nilNode {
			stack = append(stack, node)
			node = node.Left()
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, temp.Key())
			node = temp.Right()
		}
	}

	return res
}

// post-order traversal

// Recursive PostOrder, time: O(n) space: O(log n) to O(n)
func postOrderRecursiveHelper[T comparable](node, nilNode Node[T], res []T) {
	if node == nilNode {
		return
	}

	postOrderRecursiveHelper(node.Left(), nilNode, res)
	postOrderRecursiveHelper(node.Right(), nilNode, res)
	res = append(res, node.Key())
}

// Iterative PostOrder
func postOrderTwoStackHelper[T comparable](node, nilNode Node[T]) []T {
	// two stack, in first, add left then right node, on each iteration pop the top and add add to second stack
	// and repeat the operation
	var st1, st2 []Node[T]
	// add the root
	st1 = append(st1, node)
	for len(st1) > 0 {
		// pop out the top
		top := st1[len(st1)-1]
		st1 = st1[:len(st1)-1]
		// add to 2nd node
		st2 = append(st2, top)

		// if the left of top is not nil
		if top.Left() != nilNode {
			st1 = append(st1, top.Left())
		}
		// if the right node of top is not nil
		if top.Right() != nilNode {
			st1 = append(st1, top.Right())
		}
	}

	// retrieve the element and store it in result
	var res []T

	for len(st2) > 0 {
		res = append(res, st2[len(st2)-1].Key())
		st2 = st2[:len(st2)-1]
	}

	return res
}

func postOrderOneStackHelper[T comparable](node, nilNode Node[T]) []T {
	var stack []Node[T]
	var res []T
	// go left until we got right, then go left
	// if the top node's right is null, then add it to the res
	for node != nilNode || len(stack) > 0 {
		if node != nilNode {
			// add to the stack
			stack = append(stack, node)
			node = node.Left()
		} else {
			// node == nil
			temp := stack[len(stack)-1].Right()
			if temp == nilNode {
				// pop that node
				topNode := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				// add it to res
				res = append(res, topNode.Key())

				// check if it is right skewed or extreme right part (no left node),
				// below loop will execute when we are doing the operation of accessing the right part
				for len(stack) > 0 && topNode == stack[len(stack)-1].Right() {
					topNode = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					res = append(res, topNode.Key())
				}
			} else {
				node = temp
			}
		}
	}

	return res
}

// level order traversal
func levelOrderHelper[T comparable](node, nilNode Node[T]) []T {
	// using a queue iteratively
	q := queue.NewQueue[Node[T]]()
	var res []T

	q.Enqueue(node)

	for !q.Empty() {
		p := q.Dequeue()
		res = append(res, p.Key())

		if p.Left() != nilNode {
			q.Enqueue(p.Left())
		}

		if p.Right() != nilNode {
			q.Enqueue(p.Right())
		}

	}

	return res
}

func calculateDepthHelper[T comparable](node, nilNode Node[T]) int {
	if node == nilNode {
		return 0
	}

	x := calculateDepthHelper(node.Left(), nilNode)
	y := calculateDepthHelper(node.Right(), nilNode)

	if x > y {
		return x + 1
	} else {
		return y + 1
	}
}
