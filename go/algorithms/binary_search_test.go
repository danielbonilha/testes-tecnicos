package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	number := 67
	expected := 18

	result := BinarySearch(number, primes)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	number := 38
	expected := -1

	result := BinarySearch(number, primes)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}
