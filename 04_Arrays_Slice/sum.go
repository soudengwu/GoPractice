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
