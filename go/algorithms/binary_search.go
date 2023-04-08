package algorithms

import "sort"

// BinarySearch searches for a number in an ordered array
// accepts the number to be found and an ordered array
// return the index of the number, or -1 if not found
func BinarySearch(number int, arr []int) int {
	min := 0
	max := len(arr) - 1

	for {
		if max < min {
			return -1
		}

		guess := (min + max) / 2
		if arr[guess] == number {
			return guess
		}

		if arr[guess] < number {
			min = guess + 1
			continue
		}

		max = guess - 1
	}
}

// BinarySearchInt32Unsorted searches for a number in an unordered array
// accepts the number to be found and an unordered array
// return the index of the number, or -1 if not found
func BinarySearchInt32Unsorted(m []int32, number int32) int {
	sort.Slice(m, func(i, j int) bool {
		return m[i] < m[j]
	})
	i := sort.Search(len(m), func(i int) bool {
		return m[i] >= number
	})
	if i < len(m) && m[i] == number {
		return i
	}
	return -1
}
