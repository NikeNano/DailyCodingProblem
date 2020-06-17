package NodeParser

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestSerialise(t *testing.T) {
	input := node{"root", &node{"left", &node{"left.left", nil, nil}, nil}, &node{"right", nil, nil}}
	tmp, _ := serialise(input)
	output, _ := deserialise(tmp)
	fmt.Println(output.value)
	fmt.Println("hello world")
	assert.Equal(t, output.value, "left.left")
}

func TestSerialiseStrictTwo(t *testing.T) {
	input := node{"root", &node{"left", &node{"left.left", &node{}, &node{}}, &node{}}, &node{"right", &node{}, &node{}}}
	tmp := serialiseStrict(input)
	fmt.Println(tmp)
	assert.Equal(t, tmp, "hellow")

}
