package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_AddLast(t *testing.T) {

}

func TestList_AddFirst(t *testing.T) {
	tests := []struct {
		name      string
		e         int
		l         List
		expectedL List
	}{
		{
			name:      "empty list",
			e:         1,
			expectedL: List{head: &node{e: 1}, tail: &node{e: 1}, size: 1},
		},
		{
			name: "List has one element",
			e:    2,
			l:    List{head: &node{e: 1}, tail: &node{e: 1}, size: 1},
			expectedL: func() List {
				n1 := &node{e: 1}
				n2 := &node{e: 2}
				n2.next = n1
				n1.previous = n2
				return List{head: n2, tail: n1, size: 2}
			}(),
		},
		{
			name: "list has two elements",
			e:    3,
			l: func() List {
				n1 := &node{e: 1}
				n2 := &node{e: 2}
				n2.next = n1
				n1.previous = n2
				return List{head: n2, tail: n1, size: 2}
			}(),
			expectedL: func() List {
				n1 := &node{e: 1}
				n2 := &node{e: 2}
				n3 := &node{e: 3}
				n3.next = n2
				n2.previous = n3
				n2.next = n1
				n1.previous = n2
				return List{head: n3, tail: n1, size: 3}
			}(),
		},
	}

	for _, test := range tests {
		test.l.AddFirst(test.e)

		assertEqualLists(t, test.expectedL, test.l, test.name)
	}
}

func assertEqualLists(t *testing.T, expected, actual List, testName string) {
	exp, act := make([]interface{}, 0, expected.size), make([]interface{}, 0, actual.size)

	for h := expected.head; h != nil; h = h.next {
		exp = append(exp, h.e)
	}
	for h := actual.head; h != nil; h = h.next {
		act = append(act, h.e)
	}

	// TODO check correctness of previous

	assert.ElementsMatch(t, exp, act, testName)
}
