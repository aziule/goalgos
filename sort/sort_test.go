package sort

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

// List of algorithms to be tested
var sortAlgorithms = []func(elements []int) []int {
	BubbleSort,
	SelectionSort,
}

// Test all of the algorithms
func TestSort(t *testing.T) {
	expected := []int{1, 3, 4, 5, 7}

	for _, algo := range sortAlgorithms {
		elements := []int{5, 1, 4, 3, 7}
		algo(elements)
		assertEqual(t, expected, elements)
	}
}

func assertEqual(t assert.TestingT, expected, actual interface{}) {
	assert.Equal(t, expected, actual, fmt.Sprintf("Expectations do not match the actual values."), expected, actual)
}
