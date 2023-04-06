package algorithms

import (
	"math"
)

// FindCenters accepts a matrix and a radius,
// and return all possible centers of the
// matrix that fits in the radius.
//
// it is returned as a list of tuples
// ex: [[2 2] [2 3] [3 2] [3 3]]
// each tuple indicates the x and the y
// indexes respectively
func FindCenters(matrix [][]int, r int) [][]int {
	var centers [][]int

	// find the "horizontal" center (rows)
	xCenter := int(math.Ceil(float64(len(matrix))/2) - 1)

	// if the radius is bigger than the half
	// of the length, then return empty centers
	if xCenter-r < 0 {
		return centers
	}

	// find the "vertical" center (items of the rows)
	yCenter := int(math.Ceil(float64(len(matrix[xCenter]))/2) - 1)

	// if the radius is bigger than the half
	// of the length, then return empty centers
	if yCenter-r < 0 {
		return centers
	}

	// create placeholder for the indexes
	// and append the first xCenter and yCenter
	x := make([]int, 0)
	y := make([]int, 0)
	x = append(x, xCenter)
	y = append(y, yCenter)

	// scan x half up
	for i := xCenter - 1; i >= 0; i-- {
		if i-r < 0 {
			break
		}
		x = append(x, i)
	}
	// scan x half down
	for i := xCenter + 1; i < len(matrix); i++ {
		if i+r >= len(matrix) {
			break
		}
		x = append(x, i)
	}

	// scan y half left
	for i := yCenter - 1; i >= 0; i-- {
		if i-r < 0 {
			break
		}
		y = append(y, i)
	}
	// scan y half right
	for i := yCenter + 1; i < len(matrix[xCenter]); i++ {
		if i+r >= len(matrix[xCenter]) {
			break
		}
		y = append(y, i)
	}

	// build the response
	for _, v := range x {
		for _, vv := range y {
			centers = append(centers, []int{v, vv})
		}
	}

	return centers
}
