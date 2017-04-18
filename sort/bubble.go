package sort

func BubbleSort(elements []int) []int {
	nbElements := len(elements)
	swapped := false;

	for i := 0; i < nbElements - 1; i++ {
		if elements[i] > elements[i + 1] {
			elements[i], elements[i + 1] = elements[i + 1], elements[i]
			swapped = true
		}
	}

	if swapped {
		return BubbleSort(elements)
	}

	return elements
}
