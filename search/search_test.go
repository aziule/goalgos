package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestBinarySearch(t *testing.T) {
	needle := 2
	haystack := []int{1, 2, 3, 4, 5}
	expectedIndex := 1

	index, err := BinarySearch(needle, haystack)

	assert.Nil(t, err)
	assert.Equal(t, expectedIndex, index)
	assert.Equal(t, haystack[expectedIndex], needle)
}

func TestNotFoundBinarySearch(t *testing.T) {
	needle := 999
	haystack := []int{1, 2, 3, 4, 5}

	index, err := BinarySearch(needle, haystack)

	expectedError := errors.New("Needle \"999\" not found")

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, 0, index)
}
