package arrays

import "github.com/danielbonilha/testes-tecnicos/algorithms"

// MaxOccurrences given an array, count the occurrences
// of a given number and return it
//
// Example:
//
// Array: [1, 2, 2, 3, 3, 3, 3, 4, 4, 4]
// Number: 3
// Expected: 4 (occurrences)
func MaxOccurrences(arr []int, number int) int {
	index := algorithms.BinarySearch(number, arr)
	if index == -1 {
		return 0
	}

	count := 1

	// count items to the left of the index
	for i := index - 1; i >= 0; i-- {
		if arr[i] != number {
			break
		}
		count += 1
	}

	// count items to the right of the index
	for i := index + 1; i < len(arr); i++ {
		if arr[i] != number {
			break
		}
		count += 1
	}

	return count
}
