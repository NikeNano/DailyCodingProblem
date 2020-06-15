package second

import "fmt"

func compareLetters(test string, find string) bool {
	return test == find
}

func bruteFroce(word string, find string) []int {
	var index []int
	for i := range word {
		stringLenght := len(find)
		if i+stringLenght > len(word) {
			break
		}
		if compareLetters(word[i:i+stringLenght], find) {
			index = append(index, i)
		}
	}
	return index
}
