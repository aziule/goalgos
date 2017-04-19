package search

import (
	"math"
	"errors"
	"fmt"
)

func BinarySearch(needle int, haystack []int) (error, int) {
	lowerBound := 0
	upperBound := len(haystack) - 1

	return binarySearch(needle, haystack, lowerBound, upperBound)
}

// Recursive search method
func binarySearch(needle int, haystack []int, lowerBound, upperBound int) (error, int) {
	midIndex := int(math.Floor(float64((lowerBound + upperBound) / 2)))

	if lowerBound > upperBound {
		return errors.New(fmt.Sprintf("Needle \"%v\" not found", needle)), 0
	}

	if haystack[midIndex] > needle {
		return binarySearch(needle, haystack, lowerBound, midIndex - 1)
	}

	if haystack[midIndex] < needle {
		return binarySearch(needle, haystack, midIndex + 1, upperBound)
	}

	return nil, midIndex
}
