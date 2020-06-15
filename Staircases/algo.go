package first

import (
	"fmt"
)

func testy() {
	fmt.Println("Empty")
}

func StaircasesFib(n int) int {
	if n <= 1 {
		return 1
	}
	return StaircasesFib(n-1) + StaircasesFib(n-2)
}

func StaircasesItter(n int) int {
	a, b := 1, 2
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func GeneralStaircases(n int, x []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}
	var sum int
	for _, value := range x {
		sum += GeneralStaircases(n-value, x)
	}
	return sum
}

func GeneralStaircasesDynamic(n int, x []int) int {
	cache := make(map[int]int)
	cache[0] = 1
	var sum int
	for i := 1; i < n; i++ {
		for _, y := range x {
			if i-y >= 0 {
				sum += cache[i-y]
			}
		}
		cache[i] = sum
		sum = 0
	}
	return cache[n]
}
