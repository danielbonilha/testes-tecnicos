package arrays

import "testing"

func TestMaxOccurrences(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4}
	number := 3
	expected := 4

	result := MaxOccurrences(arr, number)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestMaxOccurrencesNotFound(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4}
	number := 5
	expected := 0

	result := MaxOccurrences(arr, number)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestMaxOccurrencesEdgeCase(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4}
	number := 1
	expected := 1

	result := MaxOccurrences(arr, number)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestMaxOccurrencesEmptyList(t *testing.T) {
	arr := []int{}
	number := 1
	expected := 0

	result := MaxOccurrences(arr, number)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}
