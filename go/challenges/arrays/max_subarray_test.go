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
