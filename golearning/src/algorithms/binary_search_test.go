package algorithms

import "testing"

var array = []int{1, 7, 13, 24, 35, 46, 55, 68, 79, 80, 92, 103, 114, 125, 136, 149, 150}

func TestBinarySearchRecursive_returns_correct_value_for_valid_input(t *testing.T) {
	targetValue := 55
	val := BinarySearchRecursive(array, targetValue, 0, len(array))
	if val == ELEMENT_NOT_FOUND {
		t.Fatalf("Binary Search failed to lookup value %d", targetValue)
	}
}

func TestBinarySearchRecursive_returns_failure_for_input_that_does_not_exist(t *testing.T) {
	targetValue := 3
	val := BinarySearchRecursive(array, targetValue, 0, len(array))
	if val != ELEMENT_NOT_FOUND {
		t.Fatalf("Binary Search failed to lookup value %d", targetValue)
	}
}

func TestBinarySearchRecursive_returns_failure_for_incorrect_order_of_inputs(t *testing.T) {
	targetValue := 3
	val := BinarySearchRecursive(array, targetValue, len(array), 0)
	if val != INVALID_INPUT_CONDITIONS {
		t.Fatal("Binary Search should have returned an invalid input!")
	}
}

func TestBinarySearchRecursive_returns_failure_for_empty_array(t *testing.T) {
	targetValue := 40
	val := BinarySearchRecursive([]int{}, targetValue, 0, len(array))
	if val != EMPTY_ARRAY_PASSED {
		t.Fatal("Was expecting erroneous array input")
	}
}

func TestBinarySearchRecursive_returns_failure_for_invalid_target_value(t *testing.T) {
	targetValue := -23
	val := BinarySearchRecursive(array, targetValue, len(array), 0)
	if val != INVALID_INPUT_CONDITIONS {
		t.Fatal("Binary Search should have returned an invalid input!")
	}
}

func TestBinarySearchRecursive_returns_success_for_edge_cases(t *testing.T) {
	targetValue := 1
	val := BinarySearchRecursive(array, targetValue, 0, len(array))
	if val < 0 {
		t.Fatal("Binary Search should have returned an invalid input!")
	}
}

func TestBinarySearchRecursive_returns_success_for_edge_cases_and_same_high_low_indexes(t *testing.T) {
	targetValue := 1
	val := BinarySearchRecursive(array, targetValue, 0, 0)
	if val < 0 {
		t.Fatal("Binary Search should have returned an invalid input!")
	}
}

func TestBinarySearchIterative_returns_success_for_valid_input_at_lower_end(t *testing.T) {
	targetValue := 7
	val := BinarySearchIterative(array, targetValue, 0, len(array))
	if val == ELEMENT_NOT_FOUND {
		t.Fatalf("BinarySearchIterative was supposed to find value for %d", targetValue)
	}
}

func TestBinarySearchIterative_returns_success_for_valid_input_at_upper_end(t *testing.T) {
	targetValue := 149
	val := BinarySearchIterative(array, targetValue, 0, len(array))
	if val == ELEMENT_NOT_FOUND {
		t.Fatalf("BinarySearchIterative was supposed to find value for %d", targetValue)
	}
}

func TestBinarySearchIterative_returns_failure_for_an_input_that_does_not_exist(t *testing.T) {
	targetValue := 123
	val := BinarySearchIterative(array, targetValue, 0, len(array))
	if val != ELEMENT_NOT_FOUND {
		t.Fatalf("BinarySearchIterative was supposed to fail locating the element %d", targetValue)
	}
}

func TestBinarySearchIterative_returns_failure_for_an_empty_array_input(t *testing.T) {
	targetValue := 125
	val := BinarySearchIterative([]int{}, targetValue, 0, len(array))
	if val != EMPTY_ARRAY_PASSED {
		t.Fatal("Was expecting this test to fail since an empty array was passed")
	}
}

func TestBinarySearchIterative_returns_failure_for_incorrect_order_of_input(t *testing.T) {
	targetValue := 35
	val := BinarySearchIterative(array, targetValue, len(array), 0)
	if val != INVALID_INPUT_CONDITIONS {
		t.Fatal("Was expecting this test to fail on account of incorrect ordering of high and low pointers")
	}
}

func TestBinarySearchIterative_returns_failure_for_negative_values_of_targetValue(t *testing.T) {
	targetValue := -35
	val := BinarySearchIterative(array, targetValue, 0, len(array))
	if val != ELEMENT_NOT_FOUND {
		t.Fatal("Was expecting this test to fail as targetValue is negative and does not lie in the array")
	}
}

func TestBinarySearchIterative_returns_failure_for_same_values_of_high_and_low_index_and_invalid_target_value(t *testing.T) {
	targetValue := 10
	val := BinarySearchIterative(array, targetValue, 0, 0)
	if val != ELEMENT_NOT_FOUND {
		t.Fatal("Was expecting this test to fail as targetValue is negative and does not lie in the array")
	}
}

func TestBinarySearchIterative_returns_failure_for_same_values_of_high_and_low_index_and_valid_target_value(t *testing.T) {
	targetValue := 1
	val := BinarySearchIterative(array, targetValue, 0, 0)
	if val < 0 {
		t.Fatal("This test should have returned the index of the value and not failed")
	}
}
