package algorithms

const (
	EMPTY_ARRAY_PASSED       = -999
	INVALID_INPUT_CONDITIONS = -99
	ELEMENT_NOT_FOUND        = -1
)

// BinarySearchRecursive implements the binary search in a recursive manner
func BinarySearchRecursive(numbers []int, targetVal int, lowIndex int, highIndex int) int {
	if len(numbers) == 0 {
		return EMPTY_ARRAY_PASSED
	}
	if lowIndex > highIndex || lowIndex < 0 || highIndex < 0 {
		return INVALID_INPUT_CONDITIONS
	}

	if lowIndex == highIndex && numbers[lowIndex] != targetVal {
		return ELEMENT_NOT_FOUND
	}

	mid := int((lowIndex + highIndex) / 2)
	if numbers[mid] < targetVal {
		return BinarySearchRecursive(numbers, targetVal, mid+1, highIndex)
	} else if numbers[mid] > targetVal {
		return BinarySearchRecursive(numbers, targetVal, lowIndex, mid)
	} else {
		return mid // return the index of the number
	}
}

/*
BinarySearchIterative implements the binary search in the iterative manner
*/
func BinarySearchIterative(numbers []int, targetVal int, lowIndex int, highIndex int) int {
	if len(numbers) == 0 {
		return EMPTY_ARRAY_PASSED
	}
	if lowIndex > highIndex || lowIndex < 0 || highIndex < 0 {
		return INVALID_INPUT_CONDITIONS
	}

	if lowIndex == highIndex && numbers[lowIndex] != targetVal {
		return ELEMENT_NOT_FOUND
	}

	startIndex := lowIndex
	endIndex := highIndex
	var mid int

	for startIndex <= endIndex {
		mid = int((startIndex + endIndex) / 2)

		if numbers[mid] > targetVal {
			endIndex = mid - 1
		} else if numbers[mid] < targetVal {
			startIndex = mid + 1
		} else {
			return mid
		}
	}

	return ELEMENT_NOT_FOUND
}
