package sort

func BubbleSort(elements []int) []int {
	nbElements := len(elements)
	swapped := false;

	for i := 0; i < nbElements - 1; i++ {
		if elements[i] > elements[i + 1] {
			swap(elements, i, i + 1)
			swapped = true
		}
	}

	if swapped {
		return BubbleSort(elements)
	}

	return elements
}

func swap(elements []int, x, y int) []int {
	tmp := elements[x]
	elements[x] = elements[y]
	elements[y] = tmp

	return elements
}
