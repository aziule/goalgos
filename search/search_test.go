package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

// List of algorithms to be tested
var searchAlgorithms = []func(needle int, haystack []int) (int , error) {
	BinarySearch,
	LinearSearch,
}

var haystack = []int{1, 2, 3, 4, 5}

func TestSearch(t *testing.T) {
	needle := 2
	expectedIndex := 1

	for _, algo := range searchAlgorithms {
		index, err := algo(needle, haystack)

		assert.Nil(t, err)
		assert.Equal(t, expectedIndex, index)
		assert.Equal(t, haystack[expectedIndex], needle)
	}
}

func TestSearchNotFound(t *testing.T) {
	needle := 999
	expectedError := errors.New("Needle \"999\" not found")

	for _, algo := range searchAlgorithms {
		index, err := algo(needle, haystack)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, 0, index)
	}
}
