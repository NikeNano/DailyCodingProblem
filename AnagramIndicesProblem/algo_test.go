package second

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
      
func TestGeneralStaircases(t *testing.T) {
	var answere = []int{0, 5, 8}
	assert.Equal(t, answere, bruteFroce("abcdeabcabc", "abc"))
}


// We can also use the Rabin-Karp Algo to search strings 
// Here for more info: https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
