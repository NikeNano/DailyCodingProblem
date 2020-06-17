package LowestMissingPositive

import (
	"testing"

	"gotest.tools/assert"
)

func TestSerialise(t *testing.T) {
	input := []int{3, 4, -1, 1}
	assert.Equal(t, 2, largestMissing(input))
}
