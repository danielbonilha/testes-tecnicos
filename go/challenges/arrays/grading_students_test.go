package arrays

import "testing"

func TestGradingStudentsBaseCase(t *testing.T) {
	arr := []int32{73, 67, 38, 33}
	expected := []int32{75, 67, 40, 33}

	result := GradingStudents(arr)
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("Error, expected %v, got %v", expected[i], result[i])
		}
	}
}
