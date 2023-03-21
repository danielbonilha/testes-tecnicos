package arrays

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
