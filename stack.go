package collections

import "errors"

type (
	Stack struct {
		list List
	}
)

// Push pushes an item onto the top of this stack.
func (s *Stack) Push(v interface{}) {
	s.list.AddFirst(v)
}

// Pop removes the object at the top of this stack and returns that object as the value of this function.
func (s *Stack) Pop() (interface{}, error) {
	if s.list.Size() == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.list.RemoveFirst()
}

// Peek looks at the object at the top of this stack without removing it from the stack.
func (s *Stack) Peek() (interface{}, error) {
	if s.list.Size() == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.list.GetFirst()
}

// IsEmpty tests if this stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}
