package TwoNumbersAddingUp

import (
	"testing"

	"gotest.tools/assert"
)

func TestSomething(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"
	assert.Equal(t, a, b, "The two words should be the same.")
}

func TestBruteForce(t *testing.T) {
	inputList := []int{1, 2, 3, 4}
	assert.Equal(t, true, bruteForce(inputList, 7))
	assert.Equal(t, true, bruteForce(inputList, 6))
	assert.Equal(t, false, bruteForce(inputList, 1))
	assert.Equal(t, false, bruteForce(inputList, 8))
}

func TestBruteForceDynamic(t *testing.T) {
	inputList := []int{1, 2, 3, 4}
	assert.Equal(t, true, bruteForceDynamic(inputList, 7))
	assert.Equal(t, true, bruteForceDynamic(inputList, 6))
	assert.Equal(t, false, bruteForceDynamic(inputList, 1))
	assert.Equal(t, false, bruteForceDynamic(inputList, 8))
}
