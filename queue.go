package collections

import "errors"

type (
	Queue struct {
		list List
	}
)

// Enqueue inserts the specified element at the end of this deque.
func (q *Queue) Enqueue(e interface{}) {
	q.list.AddLast(e)
}

// Dequeue retrieves and removes the first element of this deque.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.list.Size() == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.list.RemoveFirst()
}

// IsEmpty returns true if this deque contains no elements.
func (q *Queue) IsEmpty() bool {
	return q.list.Size() == 0
}
