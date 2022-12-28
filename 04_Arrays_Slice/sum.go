package arrays

// Takes in an array and returns the sum of the array
func ArraySum(numbers []int) int {
	sum := 0
	// Range returns the index and the value input
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lenghtOfNumbers := len(numbersToSum)
// 	sums := make([]int, lenghtOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = ArraySum(numbers)
// 	}

// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, ArraySum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			// slice[low:high]
			tail := numbers[1:]
			sums = append(sums, ArraySum(tail))
		}
	}

	return sums
}
