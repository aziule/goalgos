package sort

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	elements := []int{5, 1, 4, 3, 7}
	expected := []int{1, 3, 4, 5, 7}

	BubbleSort(elements)

	assert.Equal(t, expected, elements, "Expectations does not match the actual values.", expected, elements)
}
