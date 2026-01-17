package linkedlist

import "fmt"

type doublyNode struct {
	value int
	next  *doublyNode
	prev  *doublyNode
}

type DoublyLinkedList struct {
	head *doublyNode // head
	size int         // size of the ll
}

func newDoublyNode(val int) *doublyNode {
	return &doublyNode{
		value: val,
		next:  nil,
		prev:  nil,
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, size: 0}
}

// Append the node to the last
func (l *DoublyLinkedList) Append(val int) {
	node := newDoublyNode(val)
	// head at nil
	if l.head == nil {
		l.head = node
		l.size++
		return
	}

	curr := l.head
	// move to the last node
	for curr.next != nil {
		curr = curr.next
	}

	node.prev = curr
	curr.next = node
	l.size++
}

func (l *DoublyLinkedList) Insert(pos, val int) error {
	if pos < 0 || pos > l.size {
		return ErrPositionOutOfBound
	}
	// head at nil
	if l.head == nil {
		l.Append(val)
		return nil
	}

	node := newDoublyNode(val)
	// pos == 0
	if pos == 0 {
		l.head.prev = node
		node.next = l.head
		l.head = node
		l.size++
		return nil
	}

	curr := l.head

	for i := 0; i < pos-1; i++ {
		curr = curr.next
	}

	node.prev = curr
	node.next = curr.next
	if curr.next != nil {
		curr.next.prev = node
	}
	curr.next = node
	l.size++
	return nil
}

func (l *DoublyLinkedList) Get(pos int) (int, error) {
	if pos < 0 || pos >= l.size {
		return -1, ErrPositionOutOfBound
	}

	if l.head == nil {
		return -1, fmt.Errorf("head is nil")
	}

	curr := l.head
	for i := 0; i < pos; i++ {
		curr = curr.next
	}

	return curr.value, nil
}

func (l *DoublyLinkedList) Delete(pos int) (int, error) {
	if pos < 0 || pos >= l.size {
		return -1, ErrPositionOutOfBound
	}

	// head at nil
	if l.head == nil {
		return -1, fmt.Errorf("head pointing to nil")
	}

	// pos == 0
	if pos == 0 {
		temp := l.head
		l.head = temp.next

		if l.head != nil {
			l.head.prev = nil
		}

		temp.next = nil
		l.size--
		return temp.value, nil
	}

	curr := l.head
	for i := 0; i < pos; i++ {
		curr = curr.next
	}

	curr.prev.next = curr.next
	if curr.next != nil {
		curr.next.prev = curr.prev
	}
	l.size--
	return curr.value, nil
}

func (l *DoublyLinkedList) Size() int {
	return l.size
}

func (l *DoublyLinkedList) String() string {
	cur := l.head
	out := ""

	for cur != nil {
		out += fmt.Sprintf("%d <-> ", cur.value)
		cur = cur.next
	}

	out += "nil"
	return out
}
