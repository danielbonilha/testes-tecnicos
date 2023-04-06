package algorithms

import "testing"

func TestMatrixFindCentersEven(t *testing.T) {
	m := [][]int{
		{1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6},
	}
	radius := 2
	expected := [][]int{{2, 2}, {2, 3}, {3, 2}, {3, 3}}
	result := FindCenters(m, radius)

	for i, row := range expected {
		for j := range row {
			if result[i][j] != expected[i][j] {
				t.Errorf("Error, expected %v, got %v", expected, result)
			}
		}
	}
}

func TestMatrixFindCentersOdd(t *testing.T) {
	m := [][]int{
		{1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5},
	}
	radius := 2
	expected := [][]int{{2, 2}, {2, 3}}
	result := FindCenters(m, radius)

	for i, row := range expected {
		for j := range row {
			if result[i][j] != expected[i][j] {
				t.Errorf("Error, expected %v, got %v", expected, result)
			}
		}
	}
}

func TestMatrixFindCentersLargeRadius(t *testing.T) {
	m := [][]int{
		{1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5},
	}
	radius := 3
	expected := [][]int{}
	result := FindCenters(m, radius)

	if len(result) > 0 {
		t.Errorf("Error, expected %v, got %v", len(expected), len(result))
	}
}
