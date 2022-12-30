package arrays

import "fmt"

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

// we can declare array as simpleArray := [...]int{any lenght of value}
// another Array simpleArray := [...]int{5: -1, 3, 5}

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

// There will be no test for this as I am trying to understand more concept of slice
func SliceExample() {
	// numArray := [5]int{1, 2, 3, 4, 5}
	// sliceCopy := numArray[0:3]
	// fmt.Println(cap(sliceCopy[3:cap(sliceCopy)]))

	// Creating Slices.
	// You cannot increase the cap of a slice, instead you make a new slice
	var a []int = []int{1, 2, 3, 4, 5}
	a = append(a, 6)
	b := make([]int, 10)
	fmt.Println(b)

}

func SliceReverse(slice []int) []int {
	var reverse []int

	for i := len(slice) - 1; i >= 0; i-- {
		// reverse = append(reverse, slice[i])
		reverse = append(reverse, slice[i])
	}
	return reverse
}
