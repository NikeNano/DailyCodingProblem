package decodingMessages

import (
	"fmt"
	"strconv"
)

func helperDecode(input string) (string, int) {
	return input, len(input)
}

// I should pass the whole string and just cut the index ...
func numberSolutions(inputString string, stringLength int) int {
	if stringLength == 0 || stringLength == 1 {
		fmt.Println(inputString)
		return 1
	}
	count := 0
	if string(inputString[stringLength-1]) != "0" {
		fmt.Println("One")
		count = count + numberSolutions(inputString[:stringLength-1], stringLength-1)
	}
	secondInt, _ := strconv.Atoi(string(inputString[stringLength-1]))
	if string(inputString[stringLength-2]) == "2" || string(inputString[stringLength-2]) == "1" && secondInt < 7 {
		fmt.Println("Two")
		fmt.Println(inputString[:stringLength-2])
		fmt.Println(stringLength - 2)
		count = count + numberSolutions(inputString[:stringLength-1], stringLength-2)
	}
	return count
}
