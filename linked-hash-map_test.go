package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedHashSet_Add(t *testing.T) {
	l := NewLinkedHashSet()

	tests := []struct {
		name          string
		e             int
		expectedOut   bool
		expectedSize  int
		expectedFirst int
		expectedLast  int
	}{
		{
			name:          "add to empty set",
			e:             1,
			expectedOut:   true,
			expectedSize:  1,
			expectedFirst: 1,
			expectedLast:  1,
		},
		{
			name:          "add another element",
			e:             2,
			expectedOut:   true,
			expectedSize:  2,
			expectedFirst: 1,
			expectedLast:  2,
		},
		{
			name:          "add duplicate",
			e:             2,
			expectedOut:   false,
			expectedSize:  2,
			expectedFirst: 1,
			expectedLast:  2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedOut, l.Add(test.e), test.name)
		assert.Equal(t, test.expectedSize, l.Size(), test.name)
		assert.Equal(t, test.expectedFirst, *l.First(), test.name)
		assert.Equal(t, test.expectedLast, *l.Last(), test.name)
	}
}
