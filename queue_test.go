package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	actual := make([]interface{}, 0, 3)
	for !q.IsEmpty() {
		e, _ := q.Dequeue()
		actual = append(actual, e)
	}

	assert.Equal(t, []interface{}{1, 2, 3}, actual)
}

func TestQueue_Dequeue(t *testing.T) {
	q := Queue{}

	e, err := q.Dequeue()
	assert.Equal(t, nil, e)
	if assert.NotNil(t, err) {
		assert.Equal(t, "queue is empty", err.Error())
	}
}