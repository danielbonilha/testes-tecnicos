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

func TestBinarySearchNotMatch(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	number := 38
	expected := -1

	result := BinarySearch(number, primes)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestBinarySearchSingletonListMatch(t *testing.T) {
	primes := []int{2}
	number := 2
	expected := 0

	result := BinarySearch(number, primes)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestBinarySearchSingletonListNotMatch(t *testing.T) {
	primes := []int{2}
	number := 10
	expected := -1

	result := BinarySearch(number, primes)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestBinarySearchEmptyList(t *testing.T) {
	primes := make([]int, 0)
	number := 38
	expected := -1

	result := BinarySearch(number, primes)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestBinarySearchNegativeNumbersMatch(t *testing.T) {
	//primes := []int{-2, -3, -5, -7, -11, -13, -17, -19, -23, -29, -31, -37, -41, -43, -47, -53, -59, -61, -67, -71, -73, -79, -83, -89, -97}
	primes := []int{-97, -89, -83, -79, -73, -71, -67, -61, -59, -53, -47, -43, -41, -37, -31, -29, -23, -19, -17, -13, -11, -7, -5, -3, -2}
	number := -67
	expected := 6

	result := BinarySearch(number, primes)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestBinarySearchNegativeNumbersNotMatch(t *testing.T) {
	primes := []int{-97, -89, -83, -79, -73, -71, -67, -61, -59, -53, -47, -43, -41, -37, -31, -29, -23, -19, -17, -13, -11, -7, -5, -3, -2}
	number := -10
	expected := -1

	result := BinarySearch(number, primes)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}
