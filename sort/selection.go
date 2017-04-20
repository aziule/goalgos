package sort

func SelectionSort(elements []int) []int {
	nbElements := len(elements)

	for c := 0; c < nbElements; c++ {
		minIndex := c

		for i := c; i < nbElements; i++ {
			if elements[i] < elements[minIndex] {
				minIndex = i
			}
		}

		elements[c], elements[minIndex] = elements[minIndex], elements[c]
	}

	return elements
}
