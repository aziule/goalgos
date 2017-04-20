package sort

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	elements := []int{5, 1, 4, 3, 7}
	expected := []int{1, 3, 4, 5, 7}

	SelectionSort(elements)

	assert.Equal(t, expected, elements, "Expectations do not match the actual values.", expected, elements)
}
