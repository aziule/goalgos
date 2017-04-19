package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	needle := 2
	expectedIndex := 1

	haystack := []int{1, 2, 3, 4, 5}
	err, index := BinarySearch(needle, haystack)

	assert.Nil(t, err)
	assert.Equal(t, expectedIndex, index)
	assert.Equal(t, haystack[expectedIndex], needle)
}
