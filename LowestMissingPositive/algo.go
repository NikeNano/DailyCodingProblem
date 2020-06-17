package LowestMissingPositive

import "fmt"

func largestMissing(input []int) int {
	for _, value := range input {
		fmt.Println(value)
		if value < 0 { // only care about postive values
			continue
		}
		if value >= len(input) { // To large to be possible to have all the numbers up to
			continue
		}
		input[value-1] = value
	}
	if len(input) <= 1 {
		return 1
	}
	for index := range input { // This will give its order in the list, which is what we want. Place
		if input[index] != index+1 {
			return index + 1
		}
	}
	return len(input) + 1
}
