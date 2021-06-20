package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := &Stack{}
	s.Push(3)
	s.Push(2)
	s.Push(1)

	actual := make([]interface{}, 0, 3)
	for !s.IsEmpty() {
		v, _ := s.Pop()
		actual = append(actual, v)
	}

	assert.Equal(t, []interface{}{1, 2, 3}, actual)
}

func TestStack_Push(t *testing.T) {
	s := &Stack{}
	s.Push(2)
	s.Push(1)

	list := List{}
	list.AddFirst(2)
	list.AddFirst(1)
	expected := &Stack{list: list}

	assert.Equal(t, expected, s)
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name          string
		s             Stack
		expectedOut   int
		expectedErr   string
		expectedStack Stack
	}{
		{
			name:        "pop from empty stack",
			expectedErr: "stack is empty",
		},
		{
			name: "has three elements",
			s: func() Stack {
				list := List{}
				list.AddFirst(3)
				list.AddFirst(2)
				list.AddFirst(1)
				return Stack{list: list}
			}(),
			expectedOut: 1,
			expectedStack: func() Stack {
				list := List{}
				list.AddFirst(3)
				list.AddFirst(2)
				return Stack{list: list}
			}(),
		},
	}

	for _, test := range tests {
		out, err := test.s.Pop()
		assert.Equal(t, test.expectedOut, out, test.name)
		assert.Equal(t, test.expectedStack, test.s, test.name)

		if test.expectedErr == "" {
			assert.Nil(t, err, test.name)
		} else if assert.NotNil(t, err, test.name) {
			assert.Contains(t, err.Error(), test.expectedErr, test.name)
		}
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name        string
		s           Stack
		expectedOut bool
	}{
		{
			name:        "is empty",
			expectedOut: true,
		},
		{
			name: "is not empty",
			s: func() Stack {
				list := List{}
				list.AddFirst(2)
				return Stack{list: list}
			}(),
			expectedOut: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedOut, test.s.IsEmpty(), test.name)
	}
}

func TestStack_Peek(t *testing.T) {

	tests := []struct {
		name          string
		s             Stack
		expectedOut   int
		expectedErr   string
		expectedStack Stack
	}{
		{
			name:        "pop from empty stack",
			expectedErr: "stack is empty",
		},
		{
			name: "has three elements",
			s: func() Stack {
				list := List{}
				list.AddFirst(3)
				list.AddFirst(2)
				list.AddFirst(1)
				return Stack{list: list}
			}(),
			expectedOut: 1,
			expectedStack: func() Stack {
				list := List{}
				list.AddFirst(3)
				list.AddFirst(2)
				list.AddFirst(1)
				return Stack{list: list}
			}(),
		},
	}

	for _, test := range tests {
		out, err := test.s.Peek()
		assert.Equal(t, test.expectedOut, out, test.name)
		assert.Equal(t, test.expectedStack, test.s, test.name)

		if test.expectedErr == "" {
			assert.Nil(t, err, test.name)
		} else if assert.NotNil(t, err, test.name) {
			assert.Contains(t, err.Error(), test.expectedErr, test.name)
		}
	}
}
