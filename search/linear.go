package search

import (
	"errors"
	"fmt"
)

func LinearSearch(needle int, haystack []int) (int, error) {
	nbElements := len(haystack)

	for i := 0; i < nbElements; i++ {
		if haystack[i] == needle {
			return i, nil
		}
	}

	return 0, errors.New(fmt.Sprintf("Needle \"%v\" not found", needle))
}
