package arrays

import "testing"

func TestMaxSubArray(t *testing.T) {
	arr := []int{1, -2, 2, -3, -3, 3, 3, -4, 4, 4, -5}
	expected := 10

	result := MaxSubArray(arr)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestMaxSubArraySingleton(t *testing.T) {
	arr := []int{1, -4, -5}
	expected := 1

	result := MaxSubArray(arr)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}

func TestMaxSubArrayOnlyNegatives(t *testing.T) {
	arr := []int{-1, -2, -3, -4, -5}
	expected := 0

	result := MaxSubArray(arr)
	if result != expected {
		t.Errorf("Error, expected %v, got %v", expected, result)
	}
}
