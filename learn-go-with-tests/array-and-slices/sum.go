package main

func Sum(numbers []int) (sum int) {

	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {

		var sum int
		if len(numbers) == 0 {
			sum = 0
		} else {
			sum = Sum(numbers[1:])
		}
		sums = append(sums, sum)
	}
	return
}
