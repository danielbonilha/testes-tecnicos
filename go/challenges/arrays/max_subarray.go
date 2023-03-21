package arrays

// MaxSubArray given an array, find the biggest
// sum of contiguous elements and return it
//
// Example:
//
// arr := []int{1, -2, 2, -3, -3, 3, 3, -4, 4, 4, -5}
// expected := 10 (sum of 3, 3, -4, 4, 4)
func MaxSubArray(arr []int) int {
	max := 0
	current := 0

	for i := 0; i < len(arr); i++ {
		current = current + arr[i]
		if current > max {
			max = current
		}
		if current < 0 {
			current = 0
		}
	}

	return max
}
