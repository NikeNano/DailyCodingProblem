package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"
	assert.Equal(t, a, b, "The two words should be the same.")
}

func TestStaircasesFib(t *testing.T) {
	assert.Equal(t, 1, StaircasesFib(1))
	assert.Equal(t, 21, StaircasesFib(7))
}

func TestSome(t *testing.T) {
	assert.Equal(t, 21, StaircasesItter(6))
}

func TestGeneralStaircases(t *testing.T) {
	assert.Equal(t, 21, StaircasesItter(6))
}
