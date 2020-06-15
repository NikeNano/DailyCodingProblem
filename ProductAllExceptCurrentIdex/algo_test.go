package ProductAllExceptCurrentIndex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, true, true, "Hello")
}

func TestProductExcept(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	answere := []int{120, 60, 40, 30, 24}
	output := productExcept(input)
	assert.Equal(t, len(answere), len(output))
	for index := range answere {
		assert.Equal(t, answere[index], output[index])
	}
}

func TestProductExceptTwo(t *testing.T) {
	input := []int{3, 2, 1}
	answere := []int{2, 3, 6}
	output := productExcept(input)
	assert.Equal(t, len(answere), len(output))
	for index := range answere {
		assert.Equal(t, answere[index], output[index])
	}
}
