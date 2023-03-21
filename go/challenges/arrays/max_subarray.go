package arrays

func MaxSubArray(arr []int) int {
	currentSum := 0
	max := 0
	for i := 0; i < len(arr); i++ {
		max = max + arr[i]
		if currentSum < max {
			currentSum = max
		}
		if max < 0 {
			max = 0
		}
	}

	return currentSum
}
