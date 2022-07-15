package arrays

// SumReducing calculates the total from a slice of numbers.
func SumReducing(numbers []int) int {
	add := func(a, b int) int {
		return a + b
	}

	return Reduce(numbers, add, 0)
}

// SumAllTailsReducing calculates the sums of all but the first number given a collection of slices.
func SumAllTailsReducing(numbersToSum ...[]int) []int {

	sumTail := func(sums, x []int) []int {
		if len(sums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := sums[1:]
			sums = append(sums, SumReducing(tail))
		}
		return sums
	}
	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
