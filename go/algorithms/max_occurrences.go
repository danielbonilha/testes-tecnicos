package algorithms

func MaxOccurrences(arr []int, number int) int {
	index := BinarySearch(number, arr)
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

// Encontre o número de ocorrências de um
// determinado número em um array ordenado

// Exemplo:

// Array: [1, 2, 2, 3, 3, 3, 3, 4, 4, 4]
// Número: 3
// Resultado: 4 (ocorrências)
