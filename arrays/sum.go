package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nubmers := range numbersToSum {
		if len(nubmers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nubmers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
