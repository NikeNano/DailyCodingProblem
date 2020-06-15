package SmallestIntervallKsortedLists

import (
	"testing"

	"gotest.tools/assert"
)

func TestSomething(t *testing.T) {
	test := [][]int{{0, 1, 4, 17, 20, 25, 31},
		{5, 6, 10},
		{0, 3, 7, 8, 12}}
	answer := []int{3, 5}
	testAnswere := smallestValue(test)
	assert.Equal(t, len(answer), len(testAnswere))
	for index := range testAnswere {
		assert.Equal(t, answer[index], answer[index])
	}
}

func TestRemoveDuplicates(t *testing.T) {
	answere := []int{1, 2, 3}
	testInput := []int{1, 2, 3, 1, 2}
	testAnswere, err := removeDuplicates(testInput)
	if err != nil {
		assert.NilError(t, err)
	}

	assert.Equal(t, len(answere), len(testAnswere))
	for index := range testAnswere {
		assert.Equal(t, answere[index], testAnswere[index])
	}
}
