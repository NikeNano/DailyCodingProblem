package SmallestIntervallKsortedLists

import (
	"fmt"
	"sort"
	"strconv"
)

func removeDuplicates(inutList []int) ([]int, error) {
	check := make(map[string]int)
	res := make([]int, 0)

	for _, value := range inutList {
		check[strconv.Itoa(value)] = value
	}
	for value := range check {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		res = append(res, valueInt)
	}
	sort.Ints(res)
	return res, nil
}

func smallestValues(inputLists [][]int) []int {
	fmt.Println(inputLists)
	var intervall []int
	for index := range inputLists {
		intervall = append(intervall, inputLists[index][0])
	}
	interval := removeDuplicates(intervall)
	answer := []int{3, 5}
	res := []int{answer}
	for  i := interval[0]; i < interval[-1]; i++ {
		
	}
	return answer
}
