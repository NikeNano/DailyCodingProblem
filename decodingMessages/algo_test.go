package decodingMessages

import (
	"testing"

	"gotest.tools/assert"
)

func TestSerialise(t *testing.T) {
	input := "111"
	assert.Equal(t, 3, numberSolutions(helperDecode(input)))

}
