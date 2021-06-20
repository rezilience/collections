package collections

import "errors"

type (
	node struct {
		e        interface{}
		previous *node
		next     *node
	}

	List struct {
		head *node
		tail *node
		size int
	}
)

// Size returns the number of elements in this list.
func (l *List) Size() int {
	return l.size
}

// AddFirst inserts the specified element at the beginning of this List.
func (l *List) AddFirst(e interface{}) {
	n := &node{e: e}
	if l.head == nil { // List is empty
		l.head = n
		l.tail = l.head
	} else {
		l.head.previous = n
		n.next = l.head
		l.head = n
	}
	l.size++
}

// AddLast appends the specified element to the end of this List.
func (l *List) AddLast(e interface{}) {
	n := &node{e: e}
	if l.tail == nil { // List is empty
		l.head = n
		l.tail = l.head
	} else {
		l.tail.next = n
		n.previous = l.tail
		l.tail = n
	}
	l.size++
}

// RemoveFirst removes and returns the first element from this List. Returns an error if the List is empty.
func (l *List) RemoveFirst() (interface{}, error) {
	n := l.head

	switch l.size {
	case 0:
		return nil, errors.New("list is empty")
	case 1:
		l.head = nil
		l.tail = nil
	default:
		l.head.next.previous = nil
		l.head = l.head.next
		n.next = nil // TODO is this needed to avoid a memory leak?
	}
	l.size--

	return n.e, nil
}

// RemoveLast removes and returns the last element from this List.
func (l *List) RemoveLast() (interface{}, error) {
	n := l.tail

	switch l.size {
	case 0:
		return nil, errors.New("list is empty")
	case 1:
		l.head = nil
		l.tail = nil
	default:
		l.tail.previous.next = nil
		l.tail = l.tail.previous
		n.previous = nil // TODO is this needed to avoid a memory leak?
	}
	l.size--

	return n.e, nil
}

// GetFirst returns the first element in this list.
func (l *List) GetFirst() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("list is empty")
	} else {
		return l.head.e, nil
	}
}
