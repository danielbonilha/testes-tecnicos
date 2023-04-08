package arrays

import "sort"

func GradingStudents(grades []int32) []int32 {
	result := make([]int32, 0)

	// create ordered multiples
	multiples := make([]int32, 0)
	for i := int32(5); i <= 100; i += 5 {
		multiples = append(multiples, i)
	}

	// for each grade
	for _, g := range grades {
		// falling grade, append itself
		if g < 38 {
			result = append(result, g)
			continue
		}

		// search for the index in the multiples that v would fit
		r := int32(sort.Search(len(multiples), func(i int) bool {
			return multiples[int32(i)] >= g
		}))

		// apply rounding rules (negated)
		if int(r) == len(multiples) || multiples[r]-g >= 3 {
			result = append(result, g)
			continue
		}
		result = append(result, multiples[r])
	}

	return result
}
